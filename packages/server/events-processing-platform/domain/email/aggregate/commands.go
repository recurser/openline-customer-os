package aggregate

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/email/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"time"
)

func (a *EmailAggregate) CreateEmail(ctx context.Context, tenant, rawEmail, source, sourceOfTruth, appSource string, createdAt, updatedAt *time.Time) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "EmailAggregate.CreateEmail")
	defer span.Finish()
	span.LogFields(log.String("Tenant", tenant), log.String("AggregateID", a.GetID()))

	createdAtNotNil := utils.IfNotNilTimeWithDefault(createdAt, utils.Now())
	updatedAtNotNil := utils.IfNotNilTimeWithDefault(updatedAt, createdAtNotNil)
	event, err := events.NewEmailCreatedEvent(a, tenant, rawEmail, source, sourceOfTruth, appSource, createdAtNotNil, updatedAtNotNil)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewEmailCreatedEvent")
	}

	if err = event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}

func (a *EmailAggregate) UpdateEmail(ctx context.Context, tenant, sourceOfTruth string, updatedAt *time.Time) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "EmailAggregate.UpdateEmail")
	defer span.Finish()
	span.LogFields(log.String("Tenant", tenant), log.String("AggregateID", a.GetID()))

	updatedAtNotNil := utils.IfNotNilTimeWithDefault(updatedAt, utils.Now())
	if sourceOfTruth == "" {
		sourceOfTruth = a.Email.Source.SourceOfTruth
	}

	event, err := events.NewEmailUpdatedEvent(a, tenant, sourceOfTruth, updatedAtNotNil)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewEmailUpdatedEvent")
	}

	// FIXME alexb check what type of metadata should be set into event and apply it to all aggregation commands
	if err = event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}