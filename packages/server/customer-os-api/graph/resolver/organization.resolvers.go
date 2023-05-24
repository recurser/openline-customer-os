package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// OrganizationCreate is the resolver for the organization_Create field.
func (r *mutationResolver) OrganizationCreate(ctx context.Context, input model.OrganizationInput) (*model.Organization, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	createdOrganizationEntity, err := r.Services.OrganizationService.Create(ctx,
		&service.OrganizationCreateData{
			OrganizationEntity: mapper.MapOrganizationInputToEntity(&input),
			CustomFields:       mapper.MapCustomFieldInputsToEntities(input.CustomFields),
			FieldSets:          mapper.MapFieldSetInputsToEntities(input.FieldSets),
			TemplateId:         input.TemplateID,
			OrganizationTypeID: input.OrganizationTypeID,
			Domains:            input.Domains,
		})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to create organization %s", input.Name)
		return nil, err
	}
	return mapper.MapEntityToOrganization(createdOrganizationEntity), nil
}

// OrganizationUpdate is the resolver for the organization_Update field.
func (r *mutationResolver) OrganizationUpdate(ctx context.Context, input model.OrganizationUpdateInput) (*model.Organization, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	organization := mapper.MapOrganizationUpdateInputToEntity(&input)

	updatedOrganizationEntity, err := r.Services.OrganizationService.Update(ctx,
		&service.OrganizationUpdateData{
			OrganizationEntity: organization,
			OrganizationTypeID: input.OrganizationTypeID,
			Domains:            input.Domains,
		})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to update organization %s", input.ID)
		return nil, err
	}
	return mapper.MapEntityToOrganization(updatedOrganizationEntity), nil
}

// OrganizationDelete is the resolver for the organization_Delete field.
func (r *mutationResolver) OrganizationDelete(ctx context.Context, id string) (*model.Result, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	result, err := r.Services.OrganizationService.PermanentDelete(ctx, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to delete organization %s", id)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// OrganizationMerge is the resolver for the organization_Merge field.
func (r *mutationResolver) OrganizationMerge(ctx context.Context, primaryOrganizationID string, mergedOrganizationIds []string) (*model.Organization, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	for _, mergedOrganizationID := range mergedOrganizationIds {
		err := r.Services.OrganizationService.Merge(ctx, primaryOrganizationID, mergedOrganizationID)
		if err != nil {
			graphql.AddErrorf(ctx, "Failed to merge organization %s into organization %s", mergedOrganizationID, primaryOrganizationID)
			return nil, err
		}
	}

	organizationEntityPtr, err := r.Services.OrganizationService.GetOrganizationById(ctx, primaryOrganizationID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get organization by id %s", primaryOrganizationID)
		return nil, err
	}
	return mapper.MapEntityToOrganization(organizationEntityPtr), nil
}

// OrganizationAddSubsidiary is the resolver for the organization_AddSubsidiary field.
func (r *mutationResolver) OrganizationAddSubsidiary(ctx context.Context, input model.LinkOrganizationsInput) (*model.Organization, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	err := r.Services.OrganizationService.AddSubsidiary(ctx, input.OrganizationID, input.SubOrganizationID, utils.IfNotNilString(input.Type))
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to add subsidiary %s to organization %s", input.SubOrganizationID, input.OrganizationID)
		return nil, err
	}
	organizationEntity, err := r.Services.OrganizationService.GetOrganizationById(ctx, input.OrganizationID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to fetch organization %s", input.OrganizationID)
		return nil, err
	}
	return mapper.MapEntityToOrganization(organizationEntity), nil
}

// OrganizationRemoveSubsidiary is the resolver for the organization_RemoveSubsidiary field.
func (r *mutationResolver) OrganizationRemoveSubsidiary(ctx context.Context, organizationID string, subsidiaryID string) (*model.Organization, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	err := r.Services.OrganizationService.RemoveSubsidiary(ctx, organizationID, subsidiaryID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to remove subsidiary %s from organization %s", subsidiaryID, organizationID)
		return nil, err
	}
	organizationEntity, err := r.Services.OrganizationService.GetOrganizationById(ctx, organizationID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to fetch organization %s", organizationID)
		return nil, err
	}
	return mapper.MapEntityToOrganization(organizationEntity), nil
}

// OrganizationAddNewLocation is the resolver for the organization_AddNewLocation field.
func (r *mutationResolver) OrganizationAddNewLocation(ctx context.Context, organizationID string) (*model.Location, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	locationEntity, err := r.Services.LocationService.CreateLocationForEntity(ctx, entity.ORGANIZATION, organizationID, entity.SourceFields{
		Source:        entity.DataSourceOpenline,
		SourceOfTruth: entity.DataSourceOpenline,
		AppSource:     constants.AppSourceCustomerOsApi,
	})
	if err != nil {
		graphql.AddErrorf(ctx, "Error creating location for organization %s", organizationID)
		return nil, err
	}
	return mapper.MapEntityToLocation(locationEntity), nil
}

// OrganizationAddSocial is the resolver for the organization_AddSocial field.
func (r *mutationResolver) OrganizationAddSocial(ctx context.Context, organizationID string, input *model.SocialInput) (*model.Social, error) {
	panic(fmt.Errorf("not implemented: OrganizationAddSocial - organization_AddSocial"))
}

// Domains is the resolver for the domains field.
func (r *organizationResolver) Domains(ctx context.Context, obj *model.Organization) ([]string, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	domainEntities, err := dataloader.For(ctx).GetDomainsForOrganization(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get domains for organization %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToDomainNames(domainEntities), nil
}

// OrganizationType is the resolver for the organizationType field.
func (r *organizationResolver) OrganizationType(ctx context.Context, obj *model.Organization) (*model.OrganizationType, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	organizationTypeEntity, err := r.Services.OrganizationTypeService.FindOrganizationTypeForOrganization(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get organization type for organization %s", obj.ID)
		return nil, err
	}
	if organizationTypeEntity == nil {
		return nil, nil
	}
	return mapper.MapEntityToOrganizationType(organizationTypeEntity), nil
}

// Locations is the resolver for the locations field.
func (r *organizationResolver) Locations(ctx context.Context, obj *model.Organization) ([]*model.Location, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	locationEntities, err := dataloader.For(ctx).GetLocationsForOrganization(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get locations for organization %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToLocations(locationEntities), err
}

// Socials is the resolver for the socials field.
func (r *organizationResolver) Socials(ctx context.Context, obj *model.Organization) ([]*model.Social, error) {
	panic(fmt.Errorf("not implemented: Socials - socials"))
}

// Contacts is the resolver for the contacts field.
func (r *organizationResolver) Contacts(ctx context.Context, obj *model.Organization, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.ContactsPage, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.ContactService.GetContactsForOrganization(ctx, obj.ID, pagination.Page, pagination.Limit, where, sort)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not fetch contacts for organization %s", obj.ID)
		return nil, err
	}
	return &model.ContactsPage{
		Content:       mapper.MapEntitiesToContacts(paginatedResult.Rows.(*entity.ContactEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// JobRoles is the resolver for the jobRoles field.
func (r *organizationResolver) JobRoles(ctx context.Context, obj *model.Organization) ([]*model.JobRole, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	jobRoleEntities, err := dataloader.For(ctx).GetJobRolesForOrganization(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get job roles for organization %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToJobRoles(jobRoleEntities), err
}

// Notes is the resolver for the notes field.
func (r *organizationResolver) Notes(ctx context.Context, obj *model.Organization, pagination *model.Pagination) (*model.NotePage, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.NoteService.GetNotesForOrganization(ctx, obj.ID, pagination.Page, pagination.Limit)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get organization %s notes", obj.ID)
		return nil, err
	}
	return &model.NotePage{
		Content:       mapper.MapEntitiesToNotes(paginatedResult.Rows.(*entity.NoteEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// Tags is the resolver for the tags field.
func (r *organizationResolver) Tags(ctx context.Context, obj *model.Organization) ([]*model.Tag, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	tagEntities, err := dataloader.For(ctx).GetTagsForOrganization(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get tags for organization %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToTags(tagEntities), nil
}

// Emails is the resolver for the emails field.
func (r *organizationResolver) Emails(ctx context.Context, obj *model.Organization) ([]*model.Email, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	emailEntities, err := dataloader.For(ctx).GetEmailsForOrganization(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get emails for organization %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToEmails(emailEntities), nil
}

// PhoneNumbers is the resolver for the phoneNumbers field.
func (r *organizationResolver) PhoneNumbers(ctx context.Context, obj *model.Organization) ([]*model.PhoneNumber, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	phoneNumberEntities, err := dataloader.For(ctx).GetPhoneNumbersForOrganization(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get phone numbers for organization %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToPhoneNumbers(phoneNumberEntities), nil
}

// Subsidiaries is the resolver for the subsidiaries field.
func (r *organizationResolver) Subsidiaries(ctx context.Context, obj *model.Organization) ([]*model.LinkedOrganization, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	organizationEntities, err := r.Services.OrganizationService.GetSubsidiaries(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to fetch subsidiary organizations for orgnization %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToLinkedOrganizations(organizationEntities), nil
}

// SubsidiaryOf is the resolver for the subsidiaryOf field.
func (r *organizationResolver) SubsidiaryOf(ctx context.Context, obj *model.Organization) ([]*model.LinkedOrganization, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	organizationEntities, err := r.Services.OrganizationService.GetSubsidiaryOf(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to fetch subsidiary of organizations for orgnization %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToLinkedOrganizations(organizationEntities), nil
}

// CustomFields is the resolver for the customFields field.
func (r *organizationResolver) CustomFields(ctx context.Context, obj *model.Organization) ([]*model.CustomField, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	var customFields []*model.CustomField
	entityType := &model.CustomFieldEntityType{
		ID:         obj.ID,
		EntityType: model.EntityTypeOrganization,
	}
	customFieldEntities, err := r.Services.CustomFieldService.GetCustomFields(ctx, entityType)
	for _, v := range mapper.MapEntitiesToCustomFields(customFieldEntities) {
		customFields = append(customFields, v)
	}
	return customFields, err
}

// FieldSets is the resolver for the fieldSets field.
func (r *organizationResolver) FieldSets(ctx context.Context, obj *model.Organization) ([]*model.FieldSet, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	entityType := &model.CustomFieldEntityType{ID: obj.ID, EntityType: model.EntityTypeOrganization}
	fieldSetEntities, err := r.Services.FieldSetService.FindAll(ctx, entityType)
	return mapper.MapEntitiesToFieldSets(fieldSetEntities), err
}

// EntityTemplate is the resolver for the entityTemplate field.
func (r *organizationResolver) EntityTemplate(ctx context.Context, obj *model.Organization) (*model.EntityTemplate, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	entityType := &model.CustomFieldEntityType{ID: obj.ID, EntityType: model.EntityTypeOrganization}
	templateEntity, err := r.Services.EntityTemplateService.FindLinked(ctx, entityType)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get contact template for contact %s", obj.ID)
		return nil, err
	}
	if templateEntity == nil {
		return nil, nil
	}
	return mapper.MapEntityToEntityTemplate(templateEntity), err
}

// TimelineEvents is the resolver for the timelineEvents field.
func (r *organizationResolver) TimelineEvents(ctx context.Context, obj *model.Organization, from *time.Time, size int, timelineEventTypes []model.TimelineEventType) ([]model.TimelineEvent, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	timelineEvents, err := r.Services.TimelineEventService.GetTimelineEventsForOrganization(ctx, obj.ID, from, size, timelineEventTypes)
	if err != nil {
		graphql.AddErrorf(ctx, "failed to get timeline events for organization %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToTimelineEvents(timelineEvents), nil
}

// TimelineEventsTotalCount is the resolver for the timelineEventsTotalCount field.
func (r *organizationResolver) TimelineEventsTotalCount(ctx context.Context, obj *model.Organization, timelineEventTypes []model.TimelineEventType) (int64, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	count, err := r.Services.TimelineEventService.GetTimelineEventsTotalCountForOrganization(ctx, obj.ID, timelineEventTypes)
	if err != nil {
		graphql.AddErrorf(ctx, "failed to get timeline events total count for organization %s", obj.ID)
		return int64(0), err
	}
	return count, nil
}

// IssueSummaryByStatus is the resolver for the issueSummaryByStatus field.
func (r *organizationResolver) IssueSummaryByStatus(ctx context.Context, obj *model.Organization) ([]*model.IssueSummaryByStatus, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	issueCountByStatus, err := r.Services.IssueService.GetIssueSummaryByStatusForOrganization(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get issue summary by status for organization %s", obj.ID)
		return nil, err
	}
	issueSummaryByStatus := make([]*model.IssueSummaryByStatus, 0)
	for key, value := range issueCountByStatus {
		issueSummaryByStatus = append(issueSummaryByStatus, &model.IssueSummaryByStatus{
			Status: key,
			Count:  value,
		})
	}
	return issueSummaryByStatus, nil
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.OrganizationPage, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	if pagination == nil {
		pagination = &model.Pagination{Page: 0, Limit: 0}
	}
	paginatedResult, err := r.Services.OrganizationService.FindAll(ctx, pagination.Page, pagination.Limit, where, sort)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not fetch organizations")
		return nil, err
	}
	return &model.OrganizationPage{
		Content:       mapper.MapEntitiesToOrganizations(paginatedResult.Rows.(*entity.OrganizationEntities)),
		TotalPages:    paginatedResult.TotalPages,
		TotalElements: paginatedResult.TotalRows,
	}, err
}

// Organization is the resolver for the organization field.
func (r *queryResolver) Organization(ctx context.Context, id string) (*model.Organization, error) {
	defer func(start time.Time) {
		utils.LogMethodExecutionWithZap(r.log.SugarLogger(), start, utils.GetFunctionName())
	}(time.Now())

	organizationEntityPtr, err := r.Services.OrganizationService.GetOrganizationById(ctx, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get organization by id %s", id)
		return nil, err
	}
	return mapper.MapEntityToOrganization(organizationEntityPtr), nil
}

// Organization returns generated.OrganizationResolver implementation.
func (r *Resolver) Organization() generated.OrganizationResolver { return &organizationResolver{r} }

type organizationResolver struct{ *Resolver }
