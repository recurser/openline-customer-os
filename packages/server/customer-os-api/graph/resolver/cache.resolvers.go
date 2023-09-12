package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
)

// GlobalCache is the resolver for the global_Cache field.
func (r *queryResolver) GlobalCache(ctx context.Context) (*model.GlobalCache, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "GlobalCache.global_Cache", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	response := &model.GlobalCache{}

	user, err := r.Services.UserService.FindUserById(ctx, common.GetUserIdFromContext(ctx))
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed GlobalCache - find user by id")
		return nil, err
	}
	response.User = mapper.MapEntityToUser(user)

	isOwner, err := r.Services.UserService.IsOwner(ctx, user.Id)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed GlobalCache - is owner")
		return nil, err
	}
	response.IsOwner = *isOwner

	response.GCliCache = r.Services.Cache.GetStates() //pre-populate with states

	//contacts
	//for i := 'a'; i < 'z'; i++ {
	//	filter := model.Filter{}
	//	contactFirstNameStartsWith := fmt.Sprintf("%c", i)
	//	filter.Filter = &model.FilterItem{
	//		Property:      "FIRST_NAME",
	//		Operation:     model.ComparisonOperatorStartsWith,
	//		Value:         model.AnyTypeValue{Str: &contactFirstNameStartsWith},
	//		CaseSensitive: utils.BoolPtr(false),
	//	}
	//	contactsPage, err := r.Services.ContactService.FindAll(ctx, 1, 3, &filter, nil)
	//	if err != nil {
	//		tracing.TraceErr(span, err)
	//		graphql.AddErrorf(ctx, "Failed GcliCache - get contacts")
	//		return nil, err
	//	}
	//	contacts := contactsPage.Rows.(*entity.ContactEntities)
	//	for _, v := range *contacts {
	//		item := mapper.MapContactToGCliItem(v)
	//		response.GCliCache = append(response.GCliCache, &item)
	//	}
	//}

	//organizations
	//for i := 'a'; i < 'z'; i++ {
	//	filter := model.Filter{}
	//	organizationNameStartsWith := fmt.Sprintf("%c", i)
	//	filter.Filter = &model.FilterItem{
	//		Property:      "NAME",
	//		Operation:     model.ComparisonOperatorStartsWith,
	//		Value:         model.AnyTypeValue{Str: &organizationNameStartsWith},
	//		CaseSensitive: utils.BoolPtr(false),
	//	}
	//	contactsPage, err := r.Services.OrganizationService.FindAll(ctx, 1, 3, &filter, nil)
	//	if err != nil {
	//		tracing.TraceErr(span, err)
	//		graphql.AddErrorf(ctx, "Failed GcliCache - get organizations")
	//		return nil, err
	//	}
	//	organizations := contactsPage.Rows.(*entity.OrganizationEntities)
	//	for _, v := range *organizations {
	//		item := mapper.MapOrganizationToGCliItem(v)
	//		response.GCliCache = append(response.GCliCache, &item)
	//	}
	//}

	return response, nil
}
