package commands

import (
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
)

type EmailCommands struct {
	UpsertEmail UpsertEmailCommandHandler
	CreateEmail CreateEmailCommandHandler
}

func NewEmailCommands(log logger.Logger, cfg *config.Config, es eventstore.AggregateStore) *EmailCommands {
	return &EmailCommands{
		CreateEmail: NewCreateEmailCommandHandler(log, cfg, es),
		UpsertEmail: NewUpsertEmailHandler(log, cfg, es),
	}
}