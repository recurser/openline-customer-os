package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/model"
)

// DashboardViewContacts is the resolver for the dashboardView_Contacts field.
func (r *queryResolver) DashboardViewContacts(ctx context.Context, pagination model.Pagination, where *model.Filter, sort *model.SortBy) (*model.ContactsPage, error) {
	panic(fmt.Errorf("not implemented: DashboardViewContacts - dashboardView_Contacts"))
}

// DashboardViewOrganizations is the resolver for the dashboardView_Organizations field.
func (r *queryResolver) DashboardViewOrganizations(ctx context.Context, pagination model.Pagination, where *model.Filter, sort *model.SortBy) (*model.OrganizationPage, error) {
	panic(fmt.Errorf("not implemented: DashboardViewOrganizations - dashboardView_Organizations"))
}
