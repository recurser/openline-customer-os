package service

import (
	"context"
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	pb "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/user"
	cmnmod "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/common/model"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/user/command"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/user/command_handler"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/user/models"
	grpcerr "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/grpc_errors"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/opentracing/opentracing-go/log"
)

type userService struct {
	pb.UnimplementedUserGrpcServiceServer
	log          logger.Logger
	userCommands *command_handler.UserCommands
}

func NewUserService(log logger.Logger, userCommands *command_handler.UserCommands) *userService {
	return &userService{
		log:          log,
		userCommands: userCommands,
	}
}

func (s *userService) UpsertUser(ctx context.Context, request *pb.UpsertUserGrpcRequest) (*pb.UserIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "UserService.UpsertUser")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	span.LogFields(log.String("request", fmt.Sprintf("%+v", request)))

	userInputId := utils.NewUUIDIfEmpty(request.Id)

	dataFields := models.UserDataFields{
		Name:            request.Name,
		FirstName:       request.FirstName,
		LastName:        request.LastName,
		Internal:        request.Internal,
		ProfilePhotoUrl: request.ProfilePhotoUrl,
		Timezone:        request.Timezone,
	}
	sourceFields := cmnmod.Source{}
	sourceFields.FromGrpc(request.SourceFields)
	sourceFields.Source = utils.StringFirstNonEmpty(sourceFields.Source, request.Source)
	sourceFields.SourceOfTruth = utils.StringFirstNonEmpty(sourceFields.SourceOfTruth, request.SourceOfTruth)
	sourceFields.AppSource = utils.StringFirstNonEmpty(sourceFields.AppSource, request.AppSource)
	externalSystem := cmnmod.ExternalSystem{}
	externalSystem.FromGrpc(request.ExternalSystemFields)

	cmd := command.NewUpsertUserCommand(userInputId, request.Tenant, request.LoggedInUserId, sourceFields, externalSystem,
		dataFields, utils.TimestampProtoToTime(request.CreatedAt), utils.TimestampProtoToTime(request.UpdatedAt))
	if err := s.userCommands.UpsertUser.Handle(ctx, cmd); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("(UpsertUserCommand.Handle) tenant:{%s}, user input id:{%s}, err: %s", request.Tenant, userInputId, err.Error())
		return nil, s.errResponse(err)
	}

	s.log.Infof("Upserted user {%s}", userInputId)

	return &pb.UserIdGrpcResponse{Id: userInputId}, nil
}

func (s *userService) AddPlayerInfo(ctx context.Context, request *pb.AddPlayerInfoGrpcRequest) (*pb.UserIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "UserService.AddPlayerInfo")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	span.LogFields(log.String("request", fmt.Sprintf("%+v", request)))

	sourceFields := cmnmod.Source{}
	sourceFields.FromGrpc(request.SourceFields)

	cmd := command.NewAddPlayerInfoCommand(request.UserId, request.Tenant, request.LoggedInUserId, sourceFields,
		request.Provider, request.AuthId, request.IdentityId, utils.TimestampProtoToTime(request.Timestamp))
	if err := s.userCommands.AddPlayerInfo.Handle(ctx, cmd); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("(AddPlayerInfoCommand.Handle) tenant:{%s}, user input id:{%s}, err: %s", request.Tenant, request.UserId, err.Error())
		return nil, s.errResponse(err)
	}

	s.log.Infof("Added player info to user {%s}", request.UserId)

	return &pb.UserIdGrpcResponse{Id: request.UserId}, nil
}

func (s *userService) LinkJobRoleToUser(ctx context.Context, request *pb.LinkJobRoleToUserGrpcRequest) (*pb.UserIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "UserService.AddPlayerInfo")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, "")
	span.LogFields(log.String("request", fmt.Sprintf("%+v", request)))

	aggregateID := request.UserId

	cmd := command.NewLinkJobRoleCommand(aggregateID, request.Tenant, request.JobRoleId)
	if err := s.userCommands.LinkJobRoleCommand.Handle(ctx, cmd); err != nil {
		s.log.Errorf("(LinkJobRoleToUser.Handle) tenant:{%s}, user ID: {%s}, err: {%v}", request.Tenant, aggregateID, err)
		return nil, s.errResponse(err)
	}

	s.log.Infof("Linked job role {%s} to user {%s}", request.JobRoleId, aggregateID)

	return &pb.UserIdGrpcResponse{Id: aggregateID}, nil
}

func (s *userService) LinkPhoneNumberToUser(ctx context.Context, request *pb.LinkPhoneNumberToUserGrpcRequest) (*pb.UserIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "UserService.LinkPhoneNumberToUser")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	span.LogFields(log.String("request", fmt.Sprintf("%+v", request)))

	cmd := command.NewLinkPhoneNumberCommand(request.UserId, request.Tenant, request.LoggedInUserId, request.PhoneNumberId, request.Label, request.Primary)
	if err := s.userCommands.LinkPhoneNumberCommand.Handle(ctx, cmd); err != nil {
		s.log.Errorf("(LinkPhoneNumberToUser.Handle) tenant:{%s}, user ID: {%s}, err: {%v}", request.Tenant, request.UserId, err)
		return nil, s.errResponse(err)
	}

	s.log.Infof("Linked phone number {%s} to user {%s}", request.PhoneNumberId, request.UserId)

	return &pb.UserIdGrpcResponse{Id: request.UserId}, nil
}

func (s *userService) LinkEmailToUser(ctx context.Context, request *pb.LinkEmailToUserGrpcRequest) (*pb.UserIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "UserService.LinkPhoneNumberToUser")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	span.LogFields(log.String("request", fmt.Sprintf("%+v", request)))

	cmd := command.NewLinkEmailCommand(request.UserId, request.Tenant, request.LoggedInUserId, request.EmailId, request.Label, request.Primary)
	if err := s.userCommands.LinkEmailCommand.Handle(ctx, cmd); err != nil {
		s.log.Errorf("(LinkEmailToUser.Handle) tenant:{%s}, user ID: {%s}, err: {%v}", request.Tenant, request.UserId, err)
		return nil, s.errResponse(err)
	}

	s.log.Infof("Linked email {%s} to user {%s}", request.EmailId, request.UserId)

	return &pb.UserIdGrpcResponse{Id: request.UserId}, nil
}

func (s *userService) AddRole(ctx context.Context, request *pb.AddRoleGrpcRequest) (*pb.UserIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "UserService.AddRole")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	span.LogFields(log.String("request", fmt.Sprintf("%+v", request)))

	cmd := command.NewAddRole(request.UserId, request.Tenant, request.LoggedInUserId, request.Role)
	if err := s.userCommands.AddRole.Handle(ctx, cmd); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("(AddRoleCommand.Handle) tenant:{%s}, user id:{%s}, role: {%s}, err: %s", request.Tenant, request.UserId, request.Role, err.Error())
		return nil, s.errResponse(err)
	}

	s.log.Infof("Added role {%s} for user {%s}", request.Role, request.UserId)

	return &pb.UserIdGrpcResponse{Id: request.UserId}, nil
}

func (s *userService) RemoveRole(ctx context.Context, request *pb.RemoveRoleGrpcRequest) (*pb.UserIdGrpcResponse, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "UserService.RemoveRole")
	defer span.Finish()
	tracing.SetServiceSpanTags(ctx, span, request.Tenant, request.LoggedInUserId)
	span.LogFields(log.String("request", fmt.Sprintf("%+v", request)))

	cmd := command.NewRemoveRole(request.UserId, request.Tenant, request.LoggedInUserId, request.Role)
	if err := s.userCommands.RemoveRole.Handle(ctx, cmd); err != nil {
		tracing.TraceErr(span, err)
		s.log.Errorf("(RemoveRoleCommand.Handle) tenant:{%s}, user id:{%s}, role: {%s}, err: %s", request.Tenant, request.UserId, request.Role, err.Error())
		return nil, s.errResponse(err)
	}

	s.log.Infof("Removed role {%s} from user {%s}", request.Role, request.UserId)

	return &pb.UserIdGrpcResponse{Id: request.UserId}, nil
}

func (s *userService) errResponse(err error) error {
	return grpcerr.ErrResponse(err)
}