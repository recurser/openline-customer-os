package graph

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/config"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/consumers"
	contact_event_handlers "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/event_handlers"
	contact_events "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/contact/events"
	email_events "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/email/events"
	organization_event_handlers "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/event_handlers"
	organization_events "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/organization/events"
	phone_number_event_handlers "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/phone_number/event_handlers"
	phone_number_events "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/phone_number/events"
	user_event_handlers "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/user/event_handlers"
	user_events "github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/user/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"golang.org/x/sync/errgroup"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	//"golang.org/x/sync/errgroup"
)

type GraphConsumer struct {
	log                      logger.Logger
	db                       *esdb.Client
	cfg                      *config.Config
	repositories             *repository.Repositories
	phoneNumberEventHandler  *phone_number_event_handlers.GraphPhoneNumberEventHandler
	contactEventHandler      *contact_event_handlers.GraphContactEventHandler
	organizationEventHandler *organization_event_handlers.GraphOrganizationEventHandler
	emailEventHandler        *GraphEmailEventHandler
	userEventHandler         *user_event_handlers.GraphUserEventHandler
}

func NewGraphConsumer(log logger.Logger, db *esdb.Client, repositories *repository.Repositories, cfg *config.Config) *GraphConsumer {
	return &GraphConsumer{
		log:                      log,
		db:                       db,
		repositories:             repositories,
		cfg:                      cfg,
		contactEventHandler:      &contact_event_handlers.GraphContactEventHandler{Repositories: repositories},
		organizationEventHandler: &organization_event_handlers.GraphOrganizationEventHandler{Repositories: repositories},
		phoneNumberEventHandler:  &phone_number_event_handlers.GraphPhoneNumberEventHandler{Repositories: repositories},
		emailEventHandler:        &GraphEmailEventHandler{Repositories: repositories},
		userEventHandler:         &user_event_handlers.GraphUserEventHandler{Repositories: repositories},
	}
}

func (consumer *GraphConsumer) Connect(ctx context.Context, prefixes []string, poolSize int, worker consumers.Worker) error {
	consumer.subscribeToAll(ctx, prefixes)

	stream, err := consumer.db.SubscribeToPersistentSubscriptionToAll(
		ctx,
		consumer.cfg.Subscriptions.GraphSubscription.GroupName,
		esdb.SubscribeToPersistentSubscriptionOptions{},
	)
	if err != nil {
		return err
	}
	defer stream.Close()

	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i <= poolSize; i++ {
		g.Go(consumer.runWorker(ctx, worker, stream, i))
	}
	return g.Wait()
}

func (consumer *GraphConsumer) subscribeToAll(ctx context.Context, prefixes []string) {
	consumer.log.Infof("(starting graph subscription) prefixes: {%+v}", prefixes)
	settings := esdb.SubscriptionSettingsDefault()
	err := consumer.db.CreatePersistentSubscriptionToAll(ctx, consumer.cfg.Subscriptions.GraphSubscription.GroupName, esdb.PersistentAllSubscriptionOptions{
		Settings:  &settings,
		Filter:    &esdb.SubscriptionFilter{Type: esdb.StreamFilterType, Prefixes: prefixes},
		StartFrom: esdb.Start{},
	})
	if err != nil {
		if !eventstore.IsEventStoreErrorCodeResourceAlreadyExists(err) {
			consumer.log.Fatalf("(GraphConsumer.CreatePersistentSubscriptionToAll) err: {%v}", err.Error())
		} else {
			err = consumer.db.UpdatePersistentSubscriptionToAll(ctx, consumer.cfg.Subscriptions.GraphSubscription.GroupName, esdb.PersistentAllSubscriptionOptions{
				Settings: &settings,
				Filter:   &esdb.SubscriptionFilter{Type: esdb.StreamFilterType, Prefixes: prefixes},
			})
			if err != nil {
				consumer.log.Fatalf("(GraphConsumer.UpdatePersistentSubscriptionToAll) err: {%v}", err.Error())
			}
		}
	}
}

func (consumer *GraphConsumer) runWorker(ctx context.Context, worker consumers.Worker, stream *esdb.PersistentSubscription, i int) func() error {
	return func() error {
		return worker(ctx, stream, i)
	}
}

func (consumer *GraphConsumer) ProcessEvents(ctx context.Context, stream *esdb.PersistentSubscription, workerID int) error {

	for {
		event := stream.Recv()
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if event.SubscriptionDropped != nil {
			consumer.log.Errorf("(SubscriptionDropped) err: {%v}", event.SubscriptionDropped.Error)
			return errors.Wrap(event.SubscriptionDropped.Error, "Subscription Dropped")
		}

		if event.EventAppeared != nil {
			consumer.log.ConsumedEvent(constants.GraphConsumer, consumer.cfg.Subscriptions.GraphSubscription.GroupName, event.EventAppeared.Event, workerID)

			err := consumer.When(ctx, eventstore.NewEventFromRecorded(event.EventAppeared.Event.Event))
			if err != nil {
				consumer.log.Errorf("(GraphConsumer.when) err: {%v}", err)

				// FIXME alexb park event here instead of When ?  decide to retry / park etc
				if err := stream.Nack(err.Error(), esdb.NackActionRetry, event.EventAppeared.Event); err != nil {
					consumer.log.Errorf("(stream.Nack) err: {%v}", err)
					return errors.Wrap(err, "stream.Nack")
				}
			}

			err = stream.Ack(event.EventAppeared.Event)
			if err != nil {
				consumer.log.Errorf("(stream.Ack) err: {%v}", err)
				return errors.Wrap(err, "stream.Ack")
			}
			consumer.log.Infof("(ACK) event commit: {%v}", *event.EventAppeared.Event)
		}
	}
}

func (consumer *GraphConsumer) When(ctx context.Context, evt eventstore.Event) error {
	ctx, span := tracing.StartProjectionTracerSpan(ctx, "GraphConsumer.When", evt)
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()), log.String("EventType", evt.GetEventType()))

	switch evt.GetEventType() {

	case phone_number_events.PhoneNumberCreatedV1:
		return consumer.phoneNumberEventHandler.OnPhoneNumberCreate(ctx, evt)
	case phone_number_events.PhoneNumberUpdatedV1:
		return consumer.phoneNumberEventHandler.OnPhoneNumberUpdate(ctx, evt)

	case email_events.EmailCreatedV1:
		return consumer.emailEventHandler.OnEmailCreate(ctx, evt)
	case email_events.EmailUpdatedV1:
		return consumer.emailEventHandler.OnEmailUpdate(ctx, evt)
	case email_events.EmailValidationFailedV1:
		return consumer.emailEventHandler.OnEmailValidationFailed(ctx, evt)
	case email_events.EmailValidatedV1:
		return consumer.emailEventHandler.OnEmailValidated(ctx, evt)

	case contact_events.ContactCreatedV1:
		return consumer.contactEventHandler.OnContactCreate(ctx, evt)
	case contact_events.ContactUpdatedV1:
		return consumer.contactEventHandler.OnContactUpdate(ctx, evt)
	case contact_events.ContactPhoneNumberLinkedV1:
		return consumer.contactEventHandler.OnPhoneNumberLinkedToContact(ctx, evt)
	case contact_events.ContactEmailLinkedV1:
		return consumer.contactEventHandler.OnEmailLinkedToContact(ctx, evt)

	case organization_events.OrganizationCreatedV1:
		return consumer.organizationEventHandler.OnOrganizationCreate(ctx, evt)
	case organization_events.OrganizationUpdatedV1:
		return consumer.organizationEventHandler.OnOrganizationUpdate(ctx, evt)
	case organization_events.OrganizationPhoneNumberLinkedV1:
		return consumer.organizationEventHandler.OnPhoneNumberLinkedToOrganization(ctx, evt)
	case organization_events.OrganizationEmailLinkedV1:
		return consumer.organizationEventHandler.OnEmailLinkedToOrganization(ctx, evt)

	case user_events.UserCreatedV1:
		return consumer.userEventHandler.OnUserCreate(ctx, evt)
	case user_events.UserUpdatedV1:
		return consumer.userEventHandler.OnUserUpdate(ctx, evt)
	case user_events.UserPhoneNumberLinkedV1:
		return consumer.userEventHandler.OnPhoneNumberLinkedToUser(ctx, evt)
	case user_events.UserEmailLinkedV1:
		return consumer.userEventHandler.OnEmailLinkedToUser(ctx, evt)

	case "PersistentConfig1":
		consumer.log.Debugf("(GraphConsumer) [When known ignorable EventType] eventType: {%s}", evt.EventType)
		return nil

	default:
		// FIXME alexb if event was not recognized, park it
		consumer.log.Errorf("(GraphConsumer) [When unknown EventType] eventType: {%s}", evt.EventType)
		return eventstore.ErrInvalidEventType
	}
}
