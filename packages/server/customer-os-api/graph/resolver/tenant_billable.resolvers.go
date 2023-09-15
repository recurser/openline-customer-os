package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
)

// BillableInfo is the resolver for the billableInfo field.
func (r *queryResolver) BillableInfo(ctx context.Context) (*model.TenantBillableInfo, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.BillableInfo", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	result, err := r.Services.BillableService.GetBillableDetails(ctx)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to fetch billable info")
		return nil, nil
	}
	return result, err
}