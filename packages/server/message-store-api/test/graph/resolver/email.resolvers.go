package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store-api/test/graph/model"
)

// Users is the resolver for the users field.
func (r *emailResolver) Users(ctx context.Context, obj *model.Email) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Contacts is the resolver for the contacts field.
func (r *emailResolver) Contacts(ctx context.Context, obj *model.Email) ([]*model.Contact, error) {
	panic(fmt.Errorf("not implemented: Contacts - contacts"))
}

// Organizations is the resolver for the organizations field.
func (r *emailResolver) Organizations(ctx context.Context, obj *model.Email) ([]*model.Organization, error) {
	panic(fmt.Errorf("not implemented: Organizations - organizations"))
}

// EmailMergeToContact is the resolver for the emailMergeToContact field.
func (r *mutationResolver) EmailMergeToContact(ctx context.Context, contactID string, input model.EmailInput) (*model.Email, error) {
	panic(fmt.Errorf("not implemented: EmailMergeToContact - emailMergeToContact"))
}

// EmailUpdateInContact is the resolver for the emailUpdateInContact field.
func (r *mutationResolver) EmailUpdateInContact(ctx context.Context, contactID string, input model.EmailUpdateInput) (*model.Email, error) {
	panic(fmt.Errorf("not implemented: EmailUpdateInContact - emailUpdateInContact"))
}

// EmailRemoveFromContact is the resolver for the emailRemoveFromContact field.
func (r *mutationResolver) EmailRemoveFromContact(ctx context.Context, contactID string, email string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: EmailRemoveFromContact - emailRemoveFromContact"))
}

// EmailRemoveFromContactByID is the resolver for the emailRemoveFromContactById field.
func (r *mutationResolver) EmailRemoveFromContactByID(ctx context.Context, contactID string, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: EmailRemoveFromContactByID - emailRemoveFromContactById"))
}

// EmailMergeToUser is the resolver for the emailMergeToUser field.
func (r *mutationResolver) EmailMergeToUser(ctx context.Context, userID string, input model.EmailInput) (*model.Email, error) {
	panic(fmt.Errorf("not implemented: EmailMergeToUser - emailMergeToUser"))
}

// EmailUpdateInUser is the resolver for the emailUpdateInUser field.
func (r *mutationResolver) EmailUpdateInUser(ctx context.Context, userID string, input model.EmailUpdateInput) (*model.Email, error) {
	panic(fmt.Errorf("not implemented: EmailUpdateInUser - emailUpdateInUser"))
}

// EmailRemoveFromUser is the resolver for the emailRemoveFromUser field.
func (r *mutationResolver) EmailRemoveFromUser(ctx context.Context, userID string, email string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: EmailRemoveFromUser - emailRemoveFromUser"))
}

// EmailRemoveFromUserByID is the resolver for the emailRemoveFromUserById field.
func (r *mutationResolver) EmailRemoveFromUserByID(ctx context.Context, userID string, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: EmailRemoveFromUserByID - emailRemoveFromUserById"))
}

// EmailMergeToOrganization is the resolver for the emailMergeToOrganization field.
func (r *mutationResolver) EmailMergeToOrganization(ctx context.Context, organizationID string, input model.EmailInput) (*model.Email, error) {
	panic(fmt.Errorf("not implemented: EmailMergeToOrganization - emailMergeToOrganization"))
}

// EmailUpdateInOrganization is the resolver for the emailUpdateInOrganization field.
func (r *mutationResolver) EmailUpdateInOrganization(ctx context.Context, organizationID string, input model.EmailUpdateInput) (*model.Email, error) {
	panic(fmt.Errorf("not implemented: EmailUpdateInOrganization - emailUpdateInOrganization"))
}

// EmailRemoveFromOrganization is the resolver for the emailRemoveFromOrganization field.
func (r *mutationResolver) EmailRemoveFromOrganization(ctx context.Context, organizationID string, email string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: EmailRemoveFromOrganization - emailRemoveFromOrganization"))
}

// EmailRemoveFromOrganizationByID is the resolver for the emailRemoveFromOrganizationById field.
func (r *mutationResolver) EmailRemoveFromOrganizationByID(ctx context.Context, organizationID string, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: EmailRemoveFromOrganizationByID - emailRemoveFromOrganizationById"))
}

// EmailDelete is the resolver for the emailDelete field.
func (r *mutationResolver) EmailDelete(ctx context.Context, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: EmailDelete - emailDelete"))
}

// Email returns generated.EmailResolver implementation.
func (r *Resolver) Email() generated.EmailResolver { return &emailResolver{r} }

type emailResolver struct{ *Resolver }
