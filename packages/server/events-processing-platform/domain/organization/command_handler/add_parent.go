package command_handler

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/aggregate"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/command"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/validator"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type AddParentCommandHandler interface {
	Handle(ctx context.Context, command *command.AddParentCommand) error
}

type addParentCommandHandler struct {
	log logger.Logger
	es  eventstore.AggregateStore
}

func NewAddParentCommandHandler(log logger.Logger, es eventstore.AggregateStore) AddParentCommandHandler {
	return &addParentCommandHandler{log: log, es: es}
}

func (c *addParentCommandHandler) Handle(ctx context.Context, cmd *command.AddParentCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "addParentCommandHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("Tenant", cmd.Tenant), log.String("ObjectID", cmd.ObjectID))

	if err := validator.GetValidator().Struct(cmd); err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	organizationAggregate, err := aggregate.LoadOrganizationAggregate(ctx, c.es, cmd.Tenant, cmd.ObjectID)
	if err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	if err = organizationAggregate.HandleCommand(ctx, cmd); err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	return c.es.Save(ctx, organizationAggregate)
}