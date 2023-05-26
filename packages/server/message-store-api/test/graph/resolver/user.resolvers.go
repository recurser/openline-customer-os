package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/model"
)

// UserCreate is the resolver for the userCreate field.
func (r *mutationResolver) UserCreate(ctx context.Context, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserCreate - userCreate"))
}

// UserCreateInTenant is the resolver for the user_CreateInTenant field.
func (r *mutationResolver) UserCreateInTenant(ctx context.Context, input model.UserInput, tenant string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserCreateInTenant - user_CreateInTenant"))
}

// UserUpdate is the resolver for the user_Update field.
func (r *mutationResolver) UserUpdate(ctx context.Context, input model.UserUpdateInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserUpdate - user_Update"))
}

// UserAddRole is the resolver for the user_AddRole field.
func (r *mutationResolver) UserAddRole(ctx context.Context, id string, role model.Role) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserAddRole - user_AddRole"))
}

// UserRemoveRole is the resolver for the user_RemoveRole field.
func (r *mutationResolver) UserRemoveRole(ctx context.Context, id string, role model.Role) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserRemoveRole - user_RemoveRole"))
}

// UserAddRoleInTenant is the resolver for the user_AddRoleInTenant field.
func (r *mutationResolver) UserAddRoleInTenant(ctx context.Context, id string, tenant string, role model.Role) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserAddRoleInTenant - user_AddRoleInTenant"))
}

// UserRemoveRoleInTenant is the resolver for the user_RemoveRoleInTenant field.
func (r *mutationResolver) UserRemoveRoleInTenant(ctx context.Context, id string, tenant string, role model.Role) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserRemoveRoleInTenant - user_RemoveRoleInTenant"))
}

// UserDelete is the resolver for the user_Delete field.
func (r *mutationResolver) UserDelete(ctx context.Context, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: UserDelete - user_Delete"))
}

// UserDeleteInTenant is the resolver for the user_DeleteInTenant field.
func (r *mutationResolver) UserDeleteInTenant(ctx context.Context, id string, tenant string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: UserDeleteInTenant - user_DeleteInTenant"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.UserPage, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// UserByEmail is the resolver for the user_ByEmail field.
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	if r.Resolver.UserByEmail != nil {
		return r.Resolver.UserByEmail(ctx, email)
	}
	panic(fmt.Errorf("not implemented: UserByEmail - user_ByEmail"))
}

// Player is the resolver for the player field.
func (r *userResolver) Player(ctx context.Context, obj *model.User) (*model.Player, error) {
	panic(fmt.Errorf("not implemented: Player - player"))
}

// Roles is the resolver for the roles field.
func (r *userResolver) Roles(ctx context.Context, obj *model.User) ([]model.Role, error) {
	panic(fmt.Errorf("not implemented: Roles - roles"))
}

// Emails is the resolver for the emails field.
func (r *userResolver) Emails(ctx context.Context, obj *model.User) ([]*model.Email, error) {
	if r.Resolver.EmailsByUser != nil {
		return r.Resolver.EmailsByUser(ctx, obj)
	}
	panic(fmt.Errorf("not implemented: Emails - emails"))
}

// PhoneNumbers is the resolver for the phoneNumbers field.
func (r *userResolver) PhoneNumbers(ctx context.Context, obj *model.User) ([]*model.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented: PhoneNumbers - phoneNumbers"))
}

// Conversations is the resolver for the conversations field.
func (r *userResolver) Conversations(ctx context.Context, obj *model.User, pagination *model.Pagination, sort []*model.SortBy) (*model.ConversationPage, error) {
	panic(fmt.Errorf("not implemented: Conversations - conversations"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
