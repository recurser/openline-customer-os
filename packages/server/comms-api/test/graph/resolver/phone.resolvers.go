package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/model"
)

// PhoneNumberMergeToContact is the resolver for the phoneNumberMergeToContact field.
func (r *mutationResolver) PhoneNumberMergeToContact(ctx context.Context, contactID string, input model.PhoneNumberInput) (*model.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberMergeToContact - phoneNumberMergeToContact"))
}

// PhoneNumberUpdateInContact is the resolver for the phoneNumberUpdateInContact field.
func (r *mutationResolver) PhoneNumberUpdateInContact(ctx context.Context, contactID string, input model.PhoneNumberUpdateInput) (*model.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberUpdateInContact - phoneNumberUpdateInContact"))
}

// PhoneNumberRemoveFromContactByE164 is the resolver for the phoneNumberRemoveFromContactByE164 field.
func (r *mutationResolver) PhoneNumberRemoveFromContactByE164(ctx context.Context, contactID string, e164 string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberRemoveFromContactByE164 - phoneNumberRemoveFromContactByE164"))
}

// PhoneNumberRemoveFromContactByID is the resolver for the phoneNumberRemoveFromContactById field.
func (r *mutationResolver) PhoneNumberRemoveFromContactByID(ctx context.Context, contactID string, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberRemoveFromContactByID - phoneNumberRemoveFromContactById"))
}

// PhoneNumberMergeToOrganization is the resolver for the phoneNumberMergeToOrganization field.
func (r *mutationResolver) PhoneNumberMergeToOrganization(ctx context.Context, organizationID string, input model.PhoneNumberInput) (*model.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberMergeToOrganization - phoneNumberMergeToOrganization"))
}

// PhoneNumberUpdateInOrganization is the resolver for the phoneNumberUpdateInOrganization field.
func (r *mutationResolver) PhoneNumberUpdateInOrganization(ctx context.Context, organizationID string, input model.PhoneNumberUpdateInput) (*model.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberUpdateInOrganization - phoneNumberUpdateInOrganization"))
}

// PhoneNumberRemoveFromOrganizationByE164 is the resolver for the phoneNumberRemoveFromOrganizationByE164 field.
func (r *mutationResolver) PhoneNumberRemoveFromOrganizationByE164(ctx context.Context, organizationID string, e164 string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberRemoveFromOrganizationByE164 - phoneNumberRemoveFromOrganizationByE164"))
}

// PhoneNumberRemoveFromOrganizationByID is the resolver for the phoneNumberRemoveFromOrganizationById field.
func (r *mutationResolver) PhoneNumberRemoveFromOrganizationByID(ctx context.Context, organizationID string, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberRemoveFromOrganizationByID - phoneNumberRemoveFromOrganizationById"))
}

// PhoneNumberMergeToUser is the resolver for the phoneNumberMergeToUser field.
func (r *mutationResolver) PhoneNumberMergeToUser(ctx context.Context, userID string, input model.PhoneNumberInput) (*model.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberMergeToUser - phoneNumberMergeToUser"))
}

// PhoneNumberUpdateInUser is the resolver for the phoneNumberUpdateInUser field.
func (r *mutationResolver) PhoneNumberUpdateInUser(ctx context.Context, userID string, input model.PhoneNumberUpdateInput) (*model.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberUpdateInUser - phoneNumberUpdateInUser"))
}

// PhoneNumberRemoveFromUserByE164 is the resolver for the phoneNumberRemoveFromUserByE164 field.
func (r *mutationResolver) PhoneNumberRemoveFromUserByE164(ctx context.Context, userID string, e164 string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberRemoveFromUserByE164 - phoneNumberRemoveFromUserByE164"))
}

// PhoneNumberRemoveFromUserByID is the resolver for the phoneNumberRemoveFromUserById field.
func (r *mutationResolver) PhoneNumberRemoveFromUserByID(ctx context.Context, userID string, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: PhoneNumberRemoveFromUserByID - phoneNumberRemoveFromUserById"))
}

// Users is the resolver for the users field.
func (r *phoneNumberResolver) Users(ctx context.Context, obj *model.PhoneNumber) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Contacts is the resolver for the contacts field.
func (r *phoneNumberResolver) Contacts(ctx context.Context, obj *model.PhoneNumber) ([]*model.Contact, error) {
	panic(fmt.Errorf("not implemented: Contacts - contacts"))
}

// Organizations is the resolver for the organizations field.
func (r *phoneNumberResolver) Organizations(ctx context.Context, obj *model.PhoneNumber) ([]*model.Organization, error) {
	panic(fmt.Errorf("not implemented: Organizations - organizations"))
}

// PhoneNumber returns generated.PhoneNumberResolver implementation.
func (r *Resolver) PhoneNumber() generated.PhoneNumberResolver { return &phoneNumberResolver{r} }

type phoneNumberResolver struct{ *Resolver }