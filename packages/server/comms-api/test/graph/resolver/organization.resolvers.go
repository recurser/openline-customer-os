package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/model"
)

// OrganizationCreate is the resolver for the organization_Create field.
func (r *mutationResolver) OrganizationCreate(ctx context.Context, input model.OrganizationInput) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: OrganizationCreate - organization_Create"))
}

// OrganizationUpdate is the resolver for the organization_Update field.
func (r *mutationResolver) OrganizationUpdate(ctx context.Context, input model.OrganizationUpdateInput) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: OrganizationUpdate - organization_Update"))
}

// OrganizationDelete is the resolver for the organization_Delete field.
func (r *mutationResolver) OrganizationDelete(ctx context.Context, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: OrganizationDelete - organization_Delete"))
}

// OrganizationMerge is the resolver for the organization_Merge field.
func (r *mutationResolver) OrganizationMerge(ctx context.Context, primaryOrganizationID string, mergedOrganizationIds []string) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: OrganizationMerge - organization_Merge"))
}

// OrganizationAddSubsidiary is the resolver for the organization_AddSubsidiary field.
func (r *mutationResolver) OrganizationAddSubsidiary(ctx context.Context, input model.LinkOrganizationsInput) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: OrganizationAddSubsidiary - organization_AddSubsidiary"))
}

// OrganizationRemoveSubsidiary is the resolver for the organization_RemoveSubsidiary field.
func (r *mutationResolver) OrganizationRemoveSubsidiary(ctx context.Context, organizationID string, subsidiaryID string) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: OrganizationRemoveSubsidiary - organization_RemoveSubsidiary"))
}

// OrganizationAddNewLocation is the resolver for the organization_AddNewLocation field.
func (r *mutationResolver) OrganizationAddNewLocation(ctx context.Context, organizationID string) (*model.Location, error) {
	panic(fmt.Errorf("not implemented: OrganizationAddNewLocation - organization_AddNewLocation"))
}

// Domains is the resolver for the domains field.
func (r *organizationResolver) Domains(ctx context.Context, obj *model.Organization) ([]string, error) {
	panic(fmt.Errorf("not implemented: Domains - domains"))
}

// OrganizationType is the resolver for the organizationType field.
func (r *organizationResolver) OrganizationType(ctx context.Context, obj *model.Organization) (*model.OrganizationType, error) {
	panic(fmt.Errorf("not implemented: OrganizationType - organizationType"))
}

// Locations is the resolver for the locations field.
func (r *organizationResolver) Locations(ctx context.Context, obj *model.Organization) ([]*model.Location, error) {
	panic(fmt.Errorf("not implemented: Locations - locations"))
}

// Contacts is the resolver for the contacts field.
func (r *organizationResolver) Contacts(ctx context.Context, obj *model.Organization, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.ContactsPage, error) {
	panic(fmt.Errorf("not implemented: Contacts - contacts"))
}

// JobRoles is the resolver for the jobRoles field.
func (r *organizationResolver) JobRoles(ctx context.Context, obj *model.Organization) ([]*model.JobRole, error) {
	panic(fmt.Errorf("not implemented: JobRoles - jobRoles"))
}

// Notes is the resolver for the notes field.
func (r *organizationResolver) Notes(ctx context.Context, obj *model.Organization, pagination *model.Pagination) (*model.NotePage, error) {
	panic(fmt.Errorf("not implemented: Notes - notes"))
}

// Tags is the resolver for the tags field.
func (r *organizationResolver) Tags(ctx context.Context, obj *model.Organization) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented: Tags - tags"))
}

// Emails is the resolver for the emails field.
func (r *organizationResolver) Emails(ctx context.Context, obj *model.Organization) ([]*model.Email, error) {
	panic(fmt.Errorf("not implemented: Emails - emails"))
}

// PhoneNumbers is the resolver for the phoneNumbers field.
func (r *organizationResolver) PhoneNumbers(ctx context.Context, obj *model.Organization) ([]*model.PhoneNumber, error) {
	panic(fmt.Errorf("not implemented: PhoneNumbers - phoneNumbers"))
}

// Subsidiaries is the resolver for the subsidiaries field.
func (r *organizationResolver) Subsidiaries(ctx context.Context, obj *model.Organization) ([]*model.LinkedOrganization, error) {
	panic(fmt.Errorf("not implemented: Subsidiaries - subsidiaries"))
}

// SubsidiaryOf is the resolver for the subsidiaryOf field.
func (r *organizationResolver) SubsidiaryOf(ctx context.Context, obj *model.Organization) ([]*model.LinkedOrganization, error) {
	panic(fmt.Errorf("not implemented: SubsidiaryOf - subsidiaryOf"))
}

// CustomFields is the resolver for the customFields field.
func (r *organizationResolver) CustomFields(ctx context.Context, obj *model.Organization) ([]*model.CustomField, error) {
	panic(fmt.Errorf("not implemented: CustomFields - customFields"))
}

// FieldSets is the resolver for the fieldSets field.
func (r *organizationResolver) FieldSets(ctx context.Context, obj *model.Organization) ([]*model.FieldSet, error) {
	panic(fmt.Errorf("not implemented: FieldSets - fieldSets"))
}

// EntityTemplate is the resolver for the entityTemplate field.
func (r *organizationResolver) EntityTemplate(ctx context.Context, obj *model.Organization) (*model.EntityTemplate, error) {
	panic(fmt.Errorf("not implemented: EntityTemplate - entityTemplate"))
}

// TimelineEvents is the resolver for the timelineEvents field.
func (r *organizationResolver) TimelineEvents(ctx context.Context, obj *model.Organization, from *time.Time, size int, timelineEventTypes []model.TimelineEventType) ([]model.TimelineEvent, error) {
	panic(fmt.Errorf("not implemented: TimelineEvents - timelineEvents"))
}

// TimelineEventsTotalCount is the resolver for the timelineEventsTotalCount field.
func (r *organizationResolver) TimelineEventsTotalCount(ctx context.Context, obj *model.Organization, timelineEventTypes []model.TimelineEventType) (int64, error) {
	panic(fmt.Errorf("not implemented: TimelineEventsTotalCount - timelineEventsTotalCount"))
}

// IssueSummaryByStatus is the resolver for the issueSummaryByStatus field.
func (r *organizationResolver) IssueSummaryByStatus(ctx context.Context, obj *model.Organization) ([]*model.IssueSummaryByStatus, error) {
	panic(fmt.Errorf("not implemented: IssueSummaryByStatus - issueSummaryByStatus"))
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.OrganizationPage, error) {
	panic(fmt.Errorf("not implemented: Organizations - organizations"))
}

// Organization is the resolver for the organization field.
func (r *queryResolver) Organization(ctx context.Context, id string) (*model.Organization, error) {
	panic(fmt.Errorf("not implemented: Organization - organization"))
}

// Organization returns generated.OrganizationResolver implementation.
func (r *Resolver) Organization() generated.OrganizationResolver { return &organizationResolver{r} }

type organizationResolver struct{ *Resolver }
