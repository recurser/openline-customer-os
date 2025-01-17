package command_handler

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/aggregate"
	cmd "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/command"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/validator"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type UpsertOrganizationCommandHandler interface {
	Handle(ctx context.Context, command *cmd.UpsertOrganizationCommand) error
}

type upsertOrganizationCommandHandler struct {
	log logger.Logger
	es  eventstore.AggregateStore
}

func NewUpsertOrganizationCommandHandler(log logger.Logger, es eventstore.AggregateStore) UpsertOrganizationCommandHandler {
	return &upsertOrganizationCommandHandler{log: log, es: es}
}

func (c *upsertOrganizationCommandHandler) Handle(ctx context.Context, command *cmd.UpsertOrganizationCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "upsertOrganizationCommandHandler.Handle")
	defer span.Finish()
	tracing.SetCommandHandlerSpanTags(ctx, span, command.Tenant, command.LoggedInUserId)
	span.LogFields(log.String("Tenant", command.Tenant), log.String("ObjectID", command.ObjectID))

	if err := validator.GetValidator().Struct(command); err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	organizationAggregate, err := aggregate.LoadOrganizationAggregate(ctx, c.es, command.Tenant, command.ObjectID)
	if err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	orgFields := cmd.UpsertOrganizationCommandToOrganizationFieldsStruct(command)

	if aggregate.IsAggregateNotFound(organizationAggregate) {
		command.IsCreateCommand = true
		if err = organizationAggregate.CreateOrganization(ctx, orgFields, command.LoggedInUserId); err != nil {
			tracing.TraceErr(span, err)
			return err
		}
	} else {
		if err = organizationAggregate.UpdateOrganization(ctx, orgFields, command.LoggedInUserId); err != nil {
			tracing.TraceErr(span, err)
			return err
		}
	}

	return c.es.Save(ctx, organizationAggregate)
}
