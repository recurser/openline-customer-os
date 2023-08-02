package events

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/validator"
	"time"
)

const (
	InteractionEventRequestSummaryV1 = "V1_INTERACTION_EVENT_REQUEST_SUMMARY"
	InteractionEventReplaceSummaryV1 = "V1_INTERACTION_EVENT_REPLACE_SUMMARY"
)

type InteractionEventRequestSummaryEvent struct {
	Tenant      string    `json:"tenant" validate:"required"`
	RequestedAt time.Time `json:"requestedAt"`
}

func NewInteractionEventRequestSummaryEvent(aggregate eventstore.Aggregate, tenant string) (eventstore.Event, error) {
	eventData := InteractionEventRequestSummaryEvent{
		Tenant:      tenant,
		RequestedAt: utils.Now(),
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, err
	}

	event := eventstore.NewBaseEvent(aggregate, InteractionEventRequestSummaryV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, err
	}
	return event, nil
}

type InteractionEventReplaceSummaryEvent struct {
	Tenant      string    `json:"tenant" validate:"required"`
	Summary     string    `json:"summary"`
	ContentType string    `json:"contentType"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewInteractionEventReplaceSummaryEvent(aggregate eventstore.Aggregate, tenant, summary, contentType string, updatedAt time.Time) (eventstore.Event, error) {
	eventData := InteractionEventReplaceSummaryEvent{
		Tenant:      tenant,
		Summary:     summary,
		UpdatedAt:   updatedAt,
		ContentType: contentType,
	}

	if err := validator.GetValidator().Struct(eventData); err != nil {
		return eventstore.Event{}, err
	}

	event := eventstore.NewBaseEvent(aggregate, InteractionEventReplaceSummaryV1)
	if err := event.SetJsonData(&eventData); err != nil {
		return eventstore.Event{}, err
	}
	return event, nil
}