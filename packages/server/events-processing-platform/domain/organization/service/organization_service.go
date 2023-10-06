package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	pb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/organization"
	cmnmod "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/models"
	cmd "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/command"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/command_handler"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/models"
	grpcerr "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/grpc_errors"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type organizationService struct {
	pb.UnimplementedOrganizationGrpcServiceServer
	log                  logger.Logger
	repositories         *repository.Repositories
	organizationCommands *command_handler.OrganizationCommands
}

func NewOrganizationService(log logger.Logger, repositories *repository.Repositories, organizationCommands *command_handler.OrganizationCommands) *organizationService {
	return &organizationService{
		log:                  log,
		repositories:         repositories,
		organizationCommands: organizationCommands,
	}
}

func (s *organizationService) UpsertOrganization(ctx context.Context, request *pb.UpsertOrganizationGrpcRequest) (*pb.OrganizationIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "OrganizationService.UpsertOrganization")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId))

	organizationId := utils.NewUUIDIfEmpty(request.Id)

	dataFields := models.OrganizationDataFields{
		Name:              request.Name,
		Hide:              request.Hide,
		Description:       request.Description,
		Website:           request.Website,
		Industry:          request.Industry,
		SubIndustry:       request.SubIndustry,
		IndustryGroup:     request.IndustryGroup,
		TargetAudience:    request.TargetAudience,
		ValueProposition:  request.ValueProposition,
		IsPublic:          request.IsPublic,
		IsCustomer:        request.IsCustomer,
		Employees:         request.Employees,
		Market:            request.Market,
		LastFundingRound:  request.LastFundingRound,
		LastFundingAmount: request.LastFundingAmount,
		ReferenceId:       request.ReferenceId,
		Note:              request.Note,
	}
	sourceFields := cmnmod.Source{}
	sourceFields.FromGrpc(request.SourceFields)
	sourceFields.Source = utils.StringFirstNonEmpty(sourceFields.Source, request.Source)
	sourceFields.SourceOfTruth = utils.StringFirstNonEmpty(sourceFields.SourceOfTruth, request.SourceOfTruth)
	sourceFields.AppSource = utils.StringFirstNonEmpty(sourceFields.AppSource, request.AppSource)

	externalSystem := cmnmod.ExternalSystem{}
	externalSystem.FromGrpc(request.ExternalSystemFields)

	command := cmd.NewUpsertOrganizationCommand(organizationId, request.Tenant, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId),
		sourceFields, externalSystem, dataFields, utils.TimestampProtoToTime(request.CreatedAt), utils.TimestampProtoToTime(request.UpdatedAt), request.IgnoreEmptyFields)
	if err := s.organizationCommands.UpsertOrganization.Handle(ctx, command); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("(UpsertSyncOrganization.Handle) tenant:%s, organizationID: %s , err: {%v}", request.Tenant, organizationId, err)
		return nil, s.errResponse(err)
	}

	s.log.Infof("Upserted organization %s", organizationId)

	return &pb.OrganizationIdGrpcResponse{Id: organizationId}, nil
}

func (s *organizationService) LinkPhoneNumberToOrganization(ctx context.Context, request *pb.LinkPhoneNumberToOrganizationGrpcRequest) (*pb.OrganizationIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "OrganizationService.LinkPhoneNumberToOrganization")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)

	command := cmd.NewLinkPhoneNumberCommand(request.OrganizationId, request.Tenant, request.LoggedInUserId, request.PhoneNumberId, request.Label, request.Primary)
	if err := s.organizationCommands.LinkPhoneNumberCommand.Handle(ctx, command); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("(LinkPhoneNumberToOrganization.Handle) tenant:{%s}, organization ID: {%s}, err: {%v}", request.Tenant, request.OrganizationId, err)
		return nil, s.errResponse(err)
	}

	s.log.Infof("Linked phone number {%s} to organization {%s}", request.PhoneNumberId, request.OrganizationId)

	return &pb.OrganizationIdGrpcResponse{Id: request.OrganizationId}, nil
}

func (s *organizationService) LinkEmailToOrganization(ctx context.Context, request *pb.LinkEmailToOrganizationGrpcRequest) (*pb.OrganizationIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "OrganizationService.LinkEmailToOrganization")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)

	command := cmd.NewLinkEmailCommand(request.OrganizationId, request.Tenant, request.LoggedInUserId, request.EmailId, request.Label, request.Primary)
	if err := s.organizationCommands.LinkEmailCommand.Handle(ctx, command); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("tenant:{%s}, organization ID: {%s}, err: {%v}", request.Tenant, request.OrganizationId, err)
		return nil, s.errResponse(err)
	}

	s.log.Infof("Linked email {%s} to organization {%s}", request.EmailId, request.OrganizationId)

	return &pb.OrganizationIdGrpcResponse{Id: request.OrganizationId}, nil
}

func (s *organizationService) LinkDomainToOrganization(ctx context.Context, request *pb.LinkDomainToOrganizationGrpcRequest) (*pb.OrganizationIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "OrganizationService.LinkDomainToOrganization")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId))

	command := cmd.NewLinkDomainCommand(request.OrganizationId, request.Tenant, request.Domain, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId))
	if err := s.organizationCommands.LinkDomainCommand.Handle(ctx, command); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Tenant:{%s}, organization ID: {%s}, err: {%v}", request.Tenant, request.OrganizationId, err)
		return nil, s.errResponse(err)
	}

	s.log.Infof("Linked domain {%s} to organization {%s}", request.Domain, request.OrganizationId)

	return &pb.OrganizationIdGrpcResponse{Id: request.OrganizationId}, nil
}

func (s *organizationService) UpdateOrganizationRenewalLikelihood(ctx context.Context, request *pb.OrganizationRenewalLikelihoodRequest) (*pb.OrganizationIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "OrganizationService.UpdateOrganizationRenewalLikelihood")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId))
	span.LogFields(log.Object("request", request))

	// handle deadlines
	if err := ctx.Err(); err != nil {
		return nil, status.Error(codes.Canceled, "Context canceled")
	}

	fields := models.RenewalLikelihoodFields{
		RenewalLikelihood: mapper.MapRenewalLikelihoodToModels(request.Likelihood),
		Comment:           request.Comment,
		UpdatedBy:         utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId),
	}
	command := cmd.NewUpdateRenewalLikelihoodCommand(request.Tenant, request.OrganizationId, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId), fields)
	if err := s.organizationCommands.UpdateRenewalLikelihoodCommand.Handle(ctx, command); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Failed update renewal likelihood for tenant: %s organizationID: %s, err: %s", request.Tenant, request.OrganizationId, err.Error())
		return nil, s.errResponse(err)
	}

	s.log.Infof("Updated renewal likelihood for tenant:%s organizationID: %s", request.Tenant, request.OrganizationId)

	return &pb.OrganizationIdGrpcResponse{Id: request.OrganizationId}, nil
}

func (s *organizationService) UpdateOrganizationRenewalForecast(ctx context.Context, request *pb.OrganizationRenewalForecastRequest) (*pb.OrganizationIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "OrganizationService.UpdateOrganizationRenewalForecast")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId))
	span.LogFields(log.Object("request", request))

	// handle deadlines
	if err := ctx.Err(); err != nil {
		return nil, status.Error(codes.Canceled, "Context canceled")
	}

	fields := models.RenewalForecastFields{
		Amount:    request.Amount,
		Comment:   request.Comment,
		UpdatedBy: utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId),
	}
	command := cmd.NewUpdateRenewalForecastCommand(request.Tenant, request.OrganizationId, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId), fields, "")
	if err := s.organizationCommands.UpdateRenewalForecastCommand.Handle(ctx, command); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Failed update renewal forecast for tenant: %s organizationID: %s, err: %s", request.Tenant, request.OrganizationId, err.Error())
		return nil, s.errResponse(err)
	}

	s.log.Infof("Updated renewal forecast for tenant:%s organizationID: %s", request.Tenant, request.OrganizationId)

	return &pb.OrganizationIdGrpcResponse{Id: request.OrganizationId}, nil
}

func (s *organizationService) UpdateOrganizationBillingDetails(ctx context.Context, request *pb.OrganizationBillingDetailsRequest) (*pb.OrganizationIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "OrganizationService.UpdateOrganizationBillingDetails")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId))
	span.LogFields(log.Object("request", request))

	// handle deadlines
	if err := ctx.Err(); err != nil {
		return nil, status.Error(codes.Canceled, "Context canceled")
	}

	fields := models.BillingDetailsFields{
		Amount:            request.Amount,
		UpdatedBy:         utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId),
		Frequency:         mapper.MapFrequencyToString(request.Frequency),
		RenewalCycle:      mapper.MapFrequencyToString(request.RenewalCycle),
		RenewalCycleStart: utils.TimestampProtoToTime(request.CycleStart),
	}
	command := cmd.NewUpdateBillingDetailsCommand(request.Tenant, request.OrganizationId, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId), fields)
	if err := s.organizationCommands.UpdateBillingDetailsCommand.Handle(ctx, command); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Failed update billing details for tenant: %s organizationID: %s, err: %s", request.Tenant, request.OrganizationId, err.Error())
		return nil, s.errResponse(err)
	}

	s.log.Infof("Updated billing details for tenant:%s organizationID: %s", request.Tenant, request.OrganizationId)

	return &pb.OrganizationIdGrpcResponse{Id: request.OrganizationId}, nil
}

func (s *organizationService) RequestRenewNextCycleDate(ctx context.Context, request *pb.RequestRenewNextCycleDateRequest) (*pb.OrganizationIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "OrganizationService.RequestRenewNextCycleDate")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	span.LogFields(log.Object("request", request))

	// handle deadlines
	if err := ctx.Err(); err != nil {
		return nil, status.Error(codes.Canceled, "Context canceled")
	}

	command := cmd.NewRequestNextCycleDateCommand(request.Tenant, request.OrganizationId, request.LoggedInUserId)
	if err := s.organizationCommands.RequestNextCycleDateCommand.Handle(ctx, command); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Failed request next cycle date for tenant: %s organizationID: %s, err: %s", request.Tenant, request.OrganizationId, err.Error())
		return nil, s.errResponse(err)
	}

	s.log.Infof("Requested next cycle date renewal for tenant:%s organizationID: %s", request.Tenant, request.OrganizationId)

	return &pb.OrganizationIdGrpcResponse{Id: request.OrganizationId}, nil
}

func (s *organizationService) HideOrganization(ctx context.Context, request *pb.OrganizationIdGrpcRequest) (*pb.OrganizationIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "OrganizationService.HideOrganization")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId))
	span.LogFields(log.Object("request", request))

	// handle deadlines
	if err := ctx.Err(); err != nil {
		return nil, status.Error(codes.Canceled, "Context canceled")
	}

	command := cmd.NewHideOrganizationCommand(request.Tenant, request.OrganizationId, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId))
	if err := s.organizationCommands.HideOrganizationCommand.Handle(ctx, command); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Failed hide organization with id  %s for tenant %s, err: %s", request.OrganizationId, request.Tenant, err.Error())
		return nil, s.errResponse(err)
	}

	s.log.Infof("Hidden organization with id %s for tenant %s", request.OrganizationId, request.Tenant)

	return &pb.OrganizationIdGrpcResponse{Id: request.OrganizationId}, nil
}

func (s *organizationService) ShowOrganization(ctx context.Context, request *pb.OrganizationIdGrpcRequest) (*pb.OrganizationIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "OrganizationService.ShowOrganization")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId))
	span.LogFields(log.Object("request", request))

	// handle deadlines
	if err := ctx.Err(); err != nil {
		return nil, status.Error(codes.Canceled, "Context canceled")
	}

	command := cmd.NewShowOrganizationCommand(request.Tenant, request.OrganizationId, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId))
	if err := s.organizationCommands.ShowOrganizationCommand.Handle(ctx, command); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Failed show organization with id  %s for tenant %s, err: %s", request.OrganizationId, request.Tenant, err.Error())
		return nil, s.errResponse(err)
	}

	s.log.Infof("Show organization with id %s for tenant %s", request.OrganizationId, request.Tenant)

	return &pb.OrganizationIdGrpcResponse{Id: request.OrganizationId}, nil
}

func (s *organizationService) UpsertCustomFieldToOrganization(ctx context.Context, request *pb.CustomFieldForOrganizationGrpcRequest) (*pb.CustomFieldIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "OrganizationService.UpsertCustomFieldToOrganization")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId))

	customFieldId := request.CustomFieldId
	if customFieldId == "" {
		customFieldId = uuid.New().String()
	}
	sourceFields := cmnmod.Source{}
	sourceFields.FromGrpc(request.SourceFields)

	customField := models.CustomField{
		Id:         customFieldId,
		Name:       request.CustomFieldName,
		TemplateId: request.CustomFieldTemplateId,
		CustomFieldValue: models.CustomFieldValue{
			Str:     request.CustomFieldValue.StringValue,
			Bool:    request.CustomFieldValue.BoolValue,
			Time:    utils.TimestampProtoToTime(request.CustomFieldValue.DatetimeValue),
			Int:     request.CustomFieldValue.IntegerValue,
			Decimal: request.CustomFieldValue.DecimalValue,
		},
		CustomFieldDataType: mapper.MapCustomFieldDataType(request.CustomFieldDataType),
	}

	command := cmd.NewUpsertCustomFieldCommand(request.OrganizationId, request.Tenant,
		sourceFields.Source, sourceFields.SourceOfTruth, sourceFields.AppSource, utils.StringFirstNonEmpty(request.LoggedInUserId, request.UserId),
		utils.TimestampProtoToTime(request.CreatedAt), utils.TimestampProtoToTime(request.UpdatedAt), customField)
	if err := s.organizationCommands.UpsertCustomFieldCommand.Handle(ctx, command); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("Tenant:{%s}, organization ID: {%s}, err: {%v}", request.Tenant, request.OrganizationId, err)
		return nil, s.errResponse(err)
	}

	return &pb.CustomFieldIdGrpcResponse{Id: customFieldId}, nil
}

func (s *organizationService) errResponse(err error) error {
	return grpcerr.ErrResponse(err)
}
