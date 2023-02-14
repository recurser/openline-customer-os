package mapper

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/utils"
)

func MapJobRoleInputToEntity(input *model.JobRoleInput) *entity.JobRoleEntity {
	if input == nil {
		return nil
	}
	jobRoleEntity := entity.JobRoleEntity{
		JobTitle:            utils.IfNotNilString(input.JobTitle),
		Primary:             utils.IfNotNilBool(input.Primary),
		ResponsibilityLevel: utils.IfNotNilInt64(input.ResponsibilityLevel),
		Source:              entity.DataSourceOpenline,
		SourceOfTruth:       entity.DataSourceOpenline,
		AppSource:           utils.IfNotNilString(input.AppSource),
	}
	if len(jobRoleEntity.AppSource) == 0 {
		jobRoleEntity.AppSource = common.AppSourceCustomerOsApi
	}
	return &jobRoleEntity
}

func MapJobRoleUpdateInputToEntity(input *model.JobRoleUpdateInput) *entity.JobRoleEntity {
	if input == nil {
		return nil
	}
	jobRoleEntity := entity.JobRoleEntity{
		Id:                  input.ID,
		JobTitle:            utils.IfNotNilString(input.JobTitle),
		Primary:             utils.IfNotNilBool(input.Primary),
		ResponsibilityLevel: utils.IfNotNilInt64(input.ResponsibilityLevel),
		SourceOfTruth:       entity.DataSourceOpenline,
	}
	return &jobRoleEntity
}

func MapEntityToJobRole(entity *entity.JobRoleEntity) *model.JobRole {
	jobRole := model.JobRole{
		ID:                  entity.Id,
		Primary:             entity.Primary,
		Source:              MapDataSourceToModel(entity.Source),
		SourceOfTruth:       MapDataSourceToModel(entity.SourceOfTruth),
		ResponsibilityLevel: entity.ResponsibilityLevel,
		AppSource:           entity.AppSource,
		CreatedAt:           entity.CreatedAt,
		UpdatedAt:           entity.UpdatedAt,
	}
	if len(entity.JobTitle) > 0 {
		jobRole.JobTitle = utils.StringPtr(entity.JobTitle)
	}
	return &jobRole
}

func MapEntitiesToJobRoles(entities *entity.JobRoleEntities) []*model.JobRole {
	var jobRoles []*model.JobRole
	for _, jobRoleEntity := range *entities {
		jobRoles = append(jobRoles, MapEntityToJobRole(&jobRoleEntity))
	}
	return jobRoles
}