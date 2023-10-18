package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	grpccommon "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/common"
	logentrygrpc "github.com/openline-ai/openline-customer-os/packages/server/events-processing-common/gen/proto/go/api/grpc/v1/log_entry"
	"github.com/opentracing/opentracing-go/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreatedBy is the resolver for the createdBy field.
func (r *logEntryResolver) CreatedBy(ctx context.Context, obj *model.LogEntry) (*model.User, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	userEntityNillable, err := dataloader.For(ctx).GetUserAuthorForLogEntry(ctx, obj.ID)
	if err != nil {
		r.log.Errorf("Error fetching user author for log entry %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Error fetching user author for log entry %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntityToUser(userEntityNillable), nil
}

// Tags is the resolver for the tags field.
func (r *logEntryResolver) Tags(ctx context.Context, obj *model.LogEntry) ([]*model.Tag, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	tagEntities, err := dataloader.For(ctx).GetTagsForLogEntry(ctx, obj.ID)
	if err != nil {
		r.log.Errorf("Failed to get tags for log entry %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get tags for log entry %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntitiesToTags(tagEntities), nil
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *logEntryResolver) ExternalLinks(ctx context.Context, obj *model.LogEntry) ([]*model.ExternalSystem, error) {
	ctx = tracing.EnrichCtxWithSpanCtxForGraphQL(ctx, graphql.GetOperationContext(ctx))

	entities, err := dataloader.For(ctx).GetExternalSystemsForEntity(ctx, obj.ID)
	if err != nil {
		r.log.Errorf("Failed to get external systems for log entry %s: %s", obj.ID, err.Error())
		graphql.AddErrorf(ctx, "Failed to get external systems for log entry %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToExternalSystems(entities), nil
}

// LogEntryCreateForOrganization is the resolver for the logEntry_CreateForOrganization field.
func (r *mutationResolver) LogEntryCreateForOrganization(ctx context.Context, organizationID string, input model.LogEntryInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.LogEntryCreateForOrganization", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("organizationID", organizationID), log.Object("input", input))

	organizationEntity, err := r.Services.OrganizationService.GetById(ctx, organizationID)
	if err != nil || organizationEntity == nil {
		if err == nil {
			err = fmt.Errorf("organization not found")
		}
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Organization not found")
		return "", nil
	}

	response, err := r.Clients.LogEntryClient.UpsertLogEntry(ctx, &logentrygrpc.UpsertLogEntryGrpcRequest{
		Tenant:      common.GetTenantFromContext(ctx),
		UserId:      common.GetUserIdFromContext(ctx),
		Content:     utils.IfNotNilString(input.Content),
		ContentType: utils.IfNotNilString(input.ContentType),
		StartedAt:   timestamppb.New(utils.IfNotNilTimeWithDefault(input.StartedAt, utils.Now())),
		SourceFields: &grpccommon.SourceFields{
			AppSource:     constants.AppSourceCustomerOsApi,
			Source:        string(entity.DataSourceOpenline),
			SourceOfTruth: string(entity.DataSourceOpenline),
		},
		LoggedOrganizationId: utils.StringPtr(organizationID),
		AuthorUserId:         utils.StringPtr(common.GetUserIdFromContext(ctx)),
	})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Error creating log entry")
		return "", nil
	}

	for _, tag := range input.Tags {
		tagId := GetTagId(ctx, r.Services, tag.ID, tag.Name)
		if tagId == "" {
			tagEntity, _ := CreateTag(ctx, r.Services, tag.Name)
			if tagEntity != nil {
				tagId = tagEntity.Id
			}
		}
		if tagId != "" {
			_, err := r.Clients.LogEntryClient.AddTag(ctx, &logentrygrpc.AddTagGrpcRequest{
				Tenant: common.GetTenantFromContext(ctx),
				UserId: common.GetUserIdFromContext(ctx),
				Id:     response.Id,
				TagId:  tagId,
			})
			if err != nil {
				tracing.TraceErr(span, err)
				graphql.AddErrorf(ctx, "Error adding tag to log entry")
				return "", nil
			}
		}
	}
	return response.Id, nil
}

// LogEntryUpdate is the resolver for the logEntry_Update field.
func (r *mutationResolver) LogEntryUpdate(ctx context.Context, id string, input model.LogEntryUpdateInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.LogEntryUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("logEntryId", id), log.Object("input", input))

	logEntryEntity, err := r.Services.LogEntryService.GetById(ctx, id)
	if err != nil || logEntryEntity == nil {
		if err == nil {
			err = fmt.Errorf("Log entry %s not found", id)
		}
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Log entry %s not found", id)
		return "", nil
	}
	grpcRequestMessage := logentrygrpc.UpsertLogEntryGrpcRequest{
		Id:          id,
		Tenant:      common.GetTenantFromContext(ctx),
		UserId:      common.GetUserIdFromContext(ctx),
		Content:     utils.IfNotNilString(input.Content),
		ContentType: utils.IfNotNilString(input.ContentType),
		SourceFields: &grpccommon.SourceFields{
			SourceOfTruth: string(entity.DataSourceOpenline),
		},
	}
	if input.StartedAt != nil {
		grpcRequestMessage.StartedAt = timestamppb.New(*input.StartedAt)
	}

	response, err := r.Clients.LogEntryClient.UpsertLogEntry(ctx, &grpcRequestMessage)

	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Error updating log entry")
		return "", nil
	}
	return response.Id, nil
}

// LogEntryAddTag is the resolver for the logEntry_AddTag field.
func (r *mutationResolver) LogEntryAddTag(ctx context.Context, id string, input model.TagIDOrNameInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.LogEntryAddTag", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("logEntryId", id), log.Object("input", input))

	logEntryEntity, err := r.Services.LogEntryService.GetById(ctx, id)
	if err != nil || logEntryEntity == nil {
		if err == nil {
			err = fmt.Errorf("Log entry %s not found", id)
		}
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Log entry %s not found", id)
		return "", nil
	}

	tagId := GetTagId(ctx, r.Services, input.ID, input.Name)
	if tagId == "" {
		tagEntity, _ := CreateTag(ctx, r.Services, input.Name)
		if tagEntity != nil {
			tagId = tagEntity.Id
		}
	}
	if tagId != "" {
		_, err := r.Clients.LogEntryClient.AddTag(ctx, &logentrygrpc.AddTagGrpcRequest{
			Tenant: common.GetTenantFromContext(ctx),
			UserId: common.GetUserIdFromContext(ctx),
			Id:     id,
			TagId:  tagId,
		})
		if err != nil {
			tracing.TraceErr(span, err)
			graphql.AddErrorf(ctx, "Error adding tag to log entry")
			return id, nil
		}
	}
	return id, nil
}

// LogEntryRemoveTag is the resolver for the logEntry_RemoveTag field.
func (r *mutationResolver) LogEntryRemoveTag(ctx context.Context, id string, input model.TagIDOrNameInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.LogEntryAddTag", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("logEntryId", id), log.Object("input", input))

	logEntryEntity, err := r.Services.LogEntryService.GetById(ctx, id)
	if err != nil || logEntryEntity == nil {
		if err == nil {
			err = fmt.Errorf("Log entry %s not found", id)
		}
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Log entry %s not found", id)
		return "", nil
	}

	tagId := GetTagId(ctx, r.Services, input.ID, input.Name)
	if tagId != "" {
		_, err = r.Clients.LogEntryClient.RemoveTag(ctx, &logentrygrpc.RemoveTagGrpcRequest{
			Tenant: common.GetTenantFromContext(ctx),
			UserId: common.GetUserIdFromContext(ctx),
			Id:     id,
			TagId:  tagId,
		})
		if err != nil {
			tracing.TraceErr(span, err)
			graphql.AddErrorf(ctx, "Error removing tag from log entry")
			return id, nil
		}
	}
	return id, nil
}

// LogEntryResetTags is the resolver for the logEntry_ResetTags field.
func (r *mutationResolver) LogEntryResetTags(ctx context.Context, id string, input []*model.TagIDOrNameInput) (string, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.LogEntryResetTags", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("logEntryId", id), log.Object("input", input))

	logEntryEntity, err := r.Services.LogEntryService.GetById(ctx, id)
	if err != nil || logEntryEntity == nil {
		if err == nil {
			err = fmt.Errorf("Log entry %s not found", id)
		}
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Log entry %s not found", id)
		return "", nil
	}
	tags, err := r.Services.TagService.GetTagsForLogEntries(ctx, []string{logEntryEntity.Id})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Error getting tags for log entry")
		return id, nil
	}
	currentTagIds := []string{}
	for _, tag := range *tags {
		currentTagIds = append(currentTagIds, tag.Id)
	}

	newTagIds := []string{}
	for _, inputTag := range input {
		tagId := GetTagId(ctx, r.Services, inputTag.ID, inputTag.Name)
		if tagId == "" {
			tagEntity, _ := CreateTag(ctx, r.Services, inputTag.Name)
			if tagEntity != nil {
				tagId = tagEntity.Id
			}
		}
		if tagId != "" {
			newTagIds = append(newTagIds, tagId)
		}
	}

	for _, currentTagId := range currentTagIds {
		if !utils.Contains(newTagIds, currentTagId) {
			_, err = r.Clients.LogEntryClient.RemoveTag(ctx, &logentrygrpc.RemoveTagGrpcRequest{
				Tenant: common.GetTenantFromContext(ctx),
				UserId: common.GetUserIdFromContext(ctx),
				Id:     id,
				TagId:  currentTagId,
			})
			if err != nil {
				tracing.TraceErr(span, err)
				graphql.AddErrorf(ctx, "Error removing tag from log entry")
				return id, nil
			}
		}
	}
	for _, newTagId := range newTagIds {
		if !utils.Contains(currentTagIds, newTagId) {
			_, err := r.Clients.LogEntryClient.AddTag(ctx, &logentrygrpc.AddTagGrpcRequest{
				Tenant: common.GetTenantFromContext(ctx),
				UserId: common.GetUserIdFromContext(ctx),
				Id:     id,
				TagId:  newTagId,
			})
			if err != nil {
				tracing.TraceErr(span, err)
				graphql.AddErrorf(ctx, "Error adding tag to log entry")
				return id, nil
			}
		}
	}
	return id, nil
}

// LogEntry is the resolver for the logEntry field.
func (r *queryResolver) LogEntry(ctx context.Context, id string) (*model.LogEntry, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.LogEntry", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.logEntryId", id))

	logEntryEntity, err := r.Services.LogEntryService.GetById(ctx, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed getting log entry with id %s", id)
		return nil, nil
	}
	return mapper.MapEntityToLogEntry(logEntryEntity), nil
}

// LogEntry returns generated.LogEntryResolver implementation.
func (r *Resolver) LogEntry() generated.LogEntryResolver { return &logEntryResolver{r} }

type logEntryResolver struct{ *Resolver }
