package mapper

import (
	"github.com/openline-ai/openline-customer-os/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/customer-os-api/graph/model"
)

func MapEntityToContactGroup(contactGroup *entity.ContactGroupNode) *model.ContactGroup {
	return &model.ContactGroup{
		ID:   contactGroup.Id,
		Name: contactGroup.Name,
	}
}

func MapEntitiesToContactGroups(contactGroupEntities *entity.ContactGroupNodes) []*model.ContactGroup {
	var contactGroups []*model.ContactGroup
	for _, contactGroupEntity := range *contactGroupEntities {
		contactGroups = append(contactGroups, MapEntityToContactGroup(&contactGroupEntity))
	}
	return contactGroups
}
