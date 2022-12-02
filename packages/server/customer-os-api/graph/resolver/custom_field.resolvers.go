package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
)

// Definition is the resolver for the definition field.
func (r *customFieldResolver) Definition(ctx context.Context, obj *model.CustomField) (*model.CustomFieldDefinition, error) {
	entity, err := r.Services.CustomFieldDefinitionService.FindLinkedWithCustomField(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get contact definition for custom field <%s>", obj.ID)
		return nil, err
	}
	if entity == nil {
		return nil, nil
	}
	return mapper.MapEntityToCustomFieldDefinition(entity), err
}

// CustomFields is the resolver for the customFields field.
func (r *fieldSetResolver) CustomFields(ctx context.Context, obj *model.FieldSet) ([]*model.CustomField, error) {
	var customFields []*model.CustomField
	customFieldEntities, err := r.Services.CustomFieldService.FindAllForFieldSet(ctx, obj)
	for _, v := range mapper.MapEntitiesToCustomFields(customFieldEntities) {
		customFields = append(customFields, v)
	}
	return customFields, err
}

// Definition is the resolver for the definition field.
func (r *fieldSetResolver) Definition(ctx context.Context, obj *model.FieldSet) (*model.FieldSetDefinition, error) {
	entity, err := r.Services.FieldSetDefinitionService.FindLinkedWithFieldSet(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get contact definition for field set <%s>", obj.ID)
		return nil, err
	}
	if entity == nil {
		return nil, nil
	}
	return mapper.MapEntityToFieldSetDefinition(entity), err
}

// CustomFieldsMergeAndUpdateInContact is the resolver for the customFieldsMergeAndUpdateInContact field.
func (r *mutationResolver) CustomFieldsMergeAndUpdateInContact(ctx context.Context, contactID string, customFields []*model.CustomFieldInput, fieldSets []*model.FieldSetInput) (*model.Contact, error) {
	err := r.Services.CustomFieldService.MergeAndUpdateCustomFieldsForContact(ctx, contactID, mapper.MapCustomFieldInputsToEntities(customFields), mapper.MapFieldSetInputsToEntities(fieldSets))
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to merge and update custom fields for contact %s", contactID)
		return nil, err
	}
	contactEntity, err := r.Services.ContactService.FindContactById(ctx, contactID)
	if err != nil || contactEntity == nil {
		graphql.AddErrorf(ctx, "Contact with id %s not found", contactID)
		return nil, err
	}
	return mapper.MapEntityToContact(contactEntity), nil
}

// CustomFieldMergeToContact is the resolver for the customFieldMergeToContact field.
func (r *mutationResolver) CustomFieldMergeToContact(ctx context.Context, contactID string, input model.CustomFieldInput) (*model.CustomField, error) {
	result, err := r.Services.CustomFieldService.MergeCustomFieldToContact(ctx, contactID, mapper.MapCustomFieldInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not add custom field <%s> to contact <%s>", input.Name, contactID)
		return nil, err
	}
	return mapper.MapEntityToCustomField(result), nil
}

// CustomFieldUpdateInContact is the resolver for the customFieldUpdateInContact field.
func (r *mutationResolver) CustomFieldUpdateInContact(ctx context.Context, contactID string, input model.CustomFieldUpdateInput) (*model.CustomField, error) {
	result, err := r.Services.CustomFieldService.UpdateCustomFieldForContact(ctx, contactID, mapper.MapCustomFieldUpdateInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not update custom field <%s> in contact <%s>", input.ID, contactID)
		return nil, err
	}
	return mapper.MapEntityToCustomField(result), nil
}

// CustomFieldDeleteFromContactByName is the resolver for the customFieldDeleteFromContactByName field.
func (r *mutationResolver) CustomFieldDeleteFromContactByName(ctx context.Context, contactID string, fieldName string) (*model.Result, error) {
	result, err := r.Services.CustomFieldService.DeleteByNameFromContact(ctx, contactID, fieldName)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove field <%s> from contact <%s>", fieldName, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// CustomFieldDeleteFromContactByID is the resolver for the customFieldDeleteFromContactById field.
func (r *mutationResolver) CustomFieldDeleteFromContactByID(ctx context.Context, contactID string, id string) (*model.Result, error) {
	result, err := r.Services.CustomFieldService.DeleteByIdFromContact(ctx, contactID, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove custom field <%s> from contact <%s>", id, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// CustomFieldMergeToFieldSet is the resolver for the customFieldMergeToFieldSet field.
func (r *mutationResolver) CustomFieldMergeToFieldSet(ctx context.Context, contactID string, fieldSetID string, input model.CustomFieldInput) (*model.CustomField, error) {
	result, err := r.Services.CustomFieldService.MergeCustomFieldToFieldSet(ctx, contactID, fieldSetID, mapper.MapCustomFieldInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not merge custom field <%s> to contact <%s>, fields set <%s>", input.Name, contactID, fieldSetID)
		return nil, err
	}
	return mapper.MapEntityToCustomField(result), nil
}

// CustomFieldUpdateInFieldSet is the resolver for the customFieldUpdateInFieldSet field.
func (r *mutationResolver) CustomFieldUpdateInFieldSet(ctx context.Context, contactID string, fieldSetID string, input model.CustomFieldUpdateInput) (*model.CustomField, error) {
	result, err := r.Services.CustomFieldService.UpdateCustomFieldForFieldSet(ctx, contactID, fieldSetID, mapper.MapCustomFieldUpdateInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not update custom field <%s> in contact <%s>, fields set <%s>", input.ID, contactID, fieldSetID)
		return nil, err
	}
	return mapper.MapEntityToCustomField(result), nil
}

// CustomFieldDeleteFromFieldSetByID is the resolver for the customFieldDeleteFromFieldSetById field.
func (r *mutationResolver) CustomFieldDeleteFromFieldSetByID(ctx context.Context, contactID string, fieldSetID string, id string) (*model.Result, error) {
	result, err := r.Services.CustomFieldService.DeleteByIdFromFieldSet(ctx, contactID, fieldSetID, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove custom field <%s> from contact <%s>, fields set <%s>", id, contactID, fieldSetID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// FieldSetMergeToContact is the resolver for the fieldSetMergeToContact field.
func (r *mutationResolver) FieldSetMergeToContact(ctx context.Context, contactID string, input model.FieldSetInput) (*model.FieldSet, error) {
	result, err := r.Services.FieldSetService.MergeFieldSetToContact(ctx, contactID, mapper.MapFieldSetInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not merge fields set <%s> to contact <%s>", input.Name, contactID)
		return nil, err
	}
	return mapper.MapEntityToFieldSet(result), nil
}

// FieldSetUpdateInContact is the resolver for the fieldSetUpdateInContact field.
func (r *mutationResolver) FieldSetUpdateInContact(ctx context.Context, contactID string, input model.FieldSetUpdateInput) (*model.FieldSet, error) {
	result, err := r.Services.FieldSetService.UpdateFieldSetInContact(ctx, contactID, mapper.MapFieldSetUpdateInputToEntity(&input))
	if err != nil {
		graphql.AddErrorf(ctx, "Could not update fields set <%s> in contact <%s>", input.ID, contactID)
		return nil, err
	}
	return mapper.MapEntityToFieldSet(result), nil
}

// FieldSetDeleteFromContact is the resolver for the fieldSetDeleteFromContact field.
func (r *mutationResolver) FieldSetDeleteFromContact(ctx context.Context, contactID string, id string) (*model.Result, error) {
	result, err := r.Services.FieldSetService.DeleteByIdFromContact(ctx, contactID, id)
	if err != nil {
		graphql.AddErrorf(ctx, "Could not remove fields set <%s> from contact <%s>", id, contactID)
		return nil, err
	}
	return &model.Result{
		Result: result,
	}, nil
}

// CustomField returns generated.CustomFieldResolver implementation.
func (r *Resolver) CustomField() generated.CustomFieldResolver { return &customFieldResolver{r} }

// FieldSet returns generated.FieldSetResolver implementation.
func (r *Resolver) FieldSet() generated.FieldSetResolver { return &fieldSetResolver{r} }

type customFieldResolver struct{ *Resolver }
type fieldSetResolver struct{ *Resolver }
