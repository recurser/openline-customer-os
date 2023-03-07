package mapper

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

func MapEntityToInteractionSession(entity *entity.InteractionSessionEntity) *model.InteractionSession {
	return &model.InteractionSession{
		ID:            entity.Id,
		StartedAt:     entity.StartedAt,
		EndedAt:       entity.EndedAt,
		Name:          utils.StringPtr(entity.Name),
		Status:        utils.StringPtr(entity.Status),
		Type:          utils.StringPtr(entity.Type),
		Channel:       utils.StringPtr(entity.Channel),
		AppSource:     entity.AppSource,
		Source:        MapDataSourceToModel(entity.Source),
		SourceOfTruth: MapDataSourceToModel(entity.SourceOfTruth),
	}
}
