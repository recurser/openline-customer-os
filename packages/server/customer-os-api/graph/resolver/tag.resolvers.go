package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// TagCreate is the resolver for the tag_Create field.
func (r *mutationResolver) TagCreate(ctx context.Context, input model.TagInput) (*model.Tag, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.TagCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	createdTag, err := r.Services.TagService.Merge(ctx, mapper.MapTagInputToEntity(input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create tag %s", input.Name)
		return nil, err
	}
	return mapper.MapEntityToTag(*createdTag), nil
}

// TagUpdate is the resolver for the tag_Update field.
func (r *mutationResolver) TagUpdate(ctx context.Context, input model.TagUpdateInput) (*model.Tag, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.TagUpdate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.tagID", input.ID))

	updatedTag, err := r.Services.TagService.Update(ctx, mapper.MapTagUpdateInputToEntity(input))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to update tag %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToTag(*updatedTag), nil
}

// TagDelete is the resolver for the tag_Delete field.
func (r *mutationResolver) TagDelete(ctx context.Context, id string) (*model.Result, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.TagDelete", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.tagID", id))

	result, err := r.Services.TagService.UnlinkAndDelete(ctx, id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to delete tag %s", id)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context) ([]*model.Tag, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.Tags", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	tags, err := r.Services.TagService.GetAll(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to fetch tags")
		return nil, err
	}
	return mapper.MapEntitiesToTags(tags), err
}
