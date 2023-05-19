package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/model"
)

// TenantMerge is the resolver for the tenant_Merge field.
func (r *mutationResolver) TenantMerge(ctx context.Context, tenant model.TenantInput) (string, error) {
	panic(fmt.Errorf("not implemented: TenantMerge - tenant_Merge"))
}

// Tenant is the resolver for the tenant field.
func (r *queryResolver) Tenant(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: Tenant - tenant"))
}

// TenantByWorkspace is the resolver for the tenant_ByWorkspace field.
func (r *queryResolver) TenantByWorkspace(ctx context.Context, workspace model.WorkspaceInput) (*string, error) {
	panic(fmt.Errorf("not implemented: TenantByWorkspace - tenant_ByWorkspace"))
}
