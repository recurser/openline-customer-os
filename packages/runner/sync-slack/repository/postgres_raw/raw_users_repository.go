package postgres

import (
	"context"
	entity "github.com/openline-ai/openline-customer-os/packages/runner/sync-slack/entity/raw"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-slack/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

func RawUsers_AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.RawUser{})
}

func RawUsers_Save(ctx context.Context, db *gorm.DB, data string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RawRepository.RawUsers_Save")
	defer span.Finish()
	tracing.SetDefaultPostgresRepositorySpanTags(span)

	rawUser := entity.RawUser{
		Data:      data,
		EmittedAt: utils.Now(),
	}
	err := db.Save(&rawUser).Error

	if err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	return nil
}
