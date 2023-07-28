package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// InteractionSession is the resolver for the interactionSession field.
func (r *interactionEventResolver) InteractionSession(ctx context.Context, obj *model.InteractionEvent) (*model.InteractionSession, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InteractionEventResolver.InteractionSession", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionEventID", obj.ID))

	interactionSessionEntityNillable, err := dataloader.For(ctx).GetInteractionSessionForInteractionEvent(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get interaction session for interaction event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntityToInteractionSession(interactionSessionEntityNillable), nil
}

// Issue is the resolver for the issue field.
func (r *interactionEventResolver) Issue(ctx context.Context, obj *model.InteractionEvent) (*model.Issue, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InteractionEventResolver.Issue", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionEventID", obj.ID))

	issueEntityNillable, err := dataloader.For(ctx).GetIssueForInteractionEvent(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get issue for interaction event %s", obj.ID)
		return nil, nil
	}
	return mapper.MapEntityToIssue(issueEntityNillable), nil
}

// Meeting is the resolver for the meeting field.
func (r *interactionEventResolver) Meeting(ctx context.Context, obj *model.InteractionEvent) (*model.Meeting, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InteractionEventResolver.Meeting", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionEventID", obj.ID))

	meetingEntityNillable, err := dataloader.For(ctx).GetMeetingForInteractionEvent(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get meeting for interaction event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntityToMeeting(meetingEntityNillable), nil
}

// SentBy is the resolver for the sentBy field.
func (r *interactionEventResolver) SentBy(ctx context.Context, obj *model.InteractionEvent) ([]model.InteractionEventParticipant, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InteractionEventResolver.SentBy", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionEventID", obj.ID))

	participantEntities, err := dataloader.For(ctx).GetSentByParticipantsForInteractionEvent(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get participants for interaction event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToInteractionEventParticipants(participantEntities), nil
}

// SentTo is the resolver for the sentTo field.
func (r *interactionEventResolver) SentTo(ctx context.Context, obj *model.InteractionEvent) ([]model.InteractionEventParticipant, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InteractionEventResolver.SentTo", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionEventID", obj.ID))

	participantEntities, err := dataloader.For(ctx).GetSentToParticipantsForInteractionEvent(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get participants for interaction event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToInteractionEventParticipants(participantEntities), nil
}

// RepliesTo is the resolver for the repliesTo field.
func (r *interactionEventResolver) RepliesTo(ctx context.Context, obj *model.InteractionEvent) (*model.InteractionEvent, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InteractionEventResolver.RepliesTo", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionEventID", obj.ID))

	interactionEventEntities, err := dataloader.For(ctx).GetInteractionEventsForInteractionEvent(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get ReplyTo for interaction event %s", obj.ID)
		return nil, err
	}
	if len(*interactionEventEntities) > 0 {
		return mapper.MapEntityToInteractionEvent(&(*interactionEventEntities)[0]), nil
	}
	return nil, nil
}

// Includes is the resolver for the includes field.
func (r *interactionEventResolver) Includes(ctx context.Context, obj *model.InteractionEvent) ([]*model.Attachment, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InteractionEventResolver.Includes", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionEventID", obj.ID))

	entities, err := dataloader.For(ctx).GetAttachmentsForInteractionEvent(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get attachment entities for Interaction Event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToAttachment(entities), nil
}

// ActionItems is the resolver for the actionItems field.
func (r *interactionEventResolver) ActionItems(ctx context.Context, obj *model.InteractionEvent) ([]*model.ActionItem, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InteractionEventResolver.ActionItems", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionEventID", obj.ID))

	entities, err := dataloader.For(ctx).GetActionItemsForInteractionEvent(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get action items entities for Interaction Event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToActionItem(entities), nil
}

// Events is the resolver for the events field.
func (r *interactionSessionResolver) Events(ctx context.Context, obj *model.InteractionSession) ([]*model.InteractionEvent, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InteractionSessionResolver.Events", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionSessionID", obj.ID))

	interactionEventEntities, err := dataloader.For(ctx).GetInteractionEventsForInteractionSession(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get interaction events for interaction session %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToInteractionEvents(interactionEventEntities), nil
}

// AttendedBy is the resolver for the attendedBy field.
func (r *interactionSessionResolver) AttendedBy(ctx context.Context, obj *model.InteractionSession) ([]model.InteractionSessionParticipant, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InteractionSessionResolver.AttendedBy", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionSessionID", obj.ID))

	participantEntities, err := dataloader.For(ctx).GetAttendedByParticipantsForInteractionSession(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get participants for interaction event %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToInteractionSessionParticipants(participantEntities), nil
}

// Includes is the resolver for the includes field.
func (r *interactionSessionResolver) Includes(ctx context.Context, obj *model.InteractionSession) ([]*model.Attachment, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InteractionSessionResolver.Includes", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionSessionID", obj.ID))

	entities, err := dataloader.For(ctx).GetAttachmentsForInteractionSession(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get attachment entities for Interaction Session %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToAttachment(entities), nil
}

// DescribedBy is the resolver for the describedBy field.
func (r *interactionSessionResolver) DescribedBy(ctx context.Context, obj *model.InteractionSession) ([]*model.Analysis, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "InteractionSessionResolver.DescribedBy", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionSessionID", obj.ID))

	analysisEntities, err := dataloader.For(ctx).GetDescribedByForInteractionSession(ctx, obj.ID)
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to get analysis for InteractionSession %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToAnalysis(analysisEntities), nil
}

// InteractionSessionCreate is the resolver for the interactionSession_Create field.
func (r *mutationResolver) InteractionSessionCreate(ctx context.Context, session model.InteractionSessionInput) (*model.InteractionSession, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.InteractionSessionCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	interactionSessionEntity, err := r.Services.InteractionSessionService.Create(ctx,
		&service.InteractionSessionCreateData{
			InteractionSessionEntity: mapper.MapInteractionSessionInputToEntity(&session),
			AttendedBy:               service.MapInteractionSessionParticipantInputToAddressData(session.AttendedBy),
		})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create InteractionEvent")
		return nil, err
	}
	interactionEvent := mapper.MapEntityToInteractionSession(interactionSessionEntity)
	return interactionEvent, nil
}

// InteractionSessionLinkAttachment is the resolver for the interactionSession_LinkAttachment field.
func (r *mutationResolver) InteractionSessionLinkAttachment(ctx context.Context, sessionID string, attachmentID string) (*model.InteractionSession, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.InteractionSessionLinkAttachment", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.sessionID", sessionID), log.String("request.attachmentID", attachmentID))

	session, err := r.Services.InteractionSessionService.InteractionSessionLinkAttachment(ctx, sessionID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToInteractionSession(session), nil
}

// InteractionEventCreate is the resolver for the interactionEvent_Create field.
func (r *mutationResolver) InteractionEventCreate(ctx context.Context, event model.InteractionEventInput) (*model.InteractionEvent, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.InteractionEventCreate", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)

	interactionEventCreated, err := r.Services.InteractionEventService.Create(ctx, &service.InteractionEventCreateData{
		InteractionEventEntity: mapper.MapInteractionEventInputToEntity(&event),
		SessionIdentifier:      event.InteractionSession,
		MeetingIdentifier:      event.MeetingID,
		SentBy:                 service.MapInteractionEventParticipantInputToAddressData(event.SentBy),
		SentTo:                 service.MapInteractionEventParticipantInputToAddressData(event.SentTo),
		RepliesTo:              event.RepliesTo,

		Source:        entity.DataSourceOpenline,
		SourceOfTruth: entity.DataSourceOpenline,
	})
	if err != nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "Failed to create InteractionEvent")
		return nil, err
	}
	interactionEvent := mapper.MapEntityToInteractionEvent(interactionEventCreated)
	return interactionEvent, nil
}

// InteractionEventLinkAttachment is the resolver for the interactionEvent_LinkAttachment field.
func (r *mutationResolver) InteractionEventLinkAttachment(ctx context.Context, eventID string, attachmentID string) (*model.InteractionEvent, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "MutationResolver.InteractionEventLinkAttachment", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.eventID", eventID), log.String("request.attachmentID", attachmentID))

	event, err := r.Services.InteractionEventService.InteractionEventLinkAttachment(ctx, eventID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToInteractionEvent(event), nil
}

// InteractionSession is the resolver for the interactionSession field.
func (r *queryResolver) InteractionSession(ctx context.Context, id string) (*model.InteractionSession, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.InteractionSession", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionSessionID", id))

	interactionSessionEntity, err := r.Services.InteractionSessionService.GetInteractionSessionById(ctx, id)
	if err != nil || interactionSessionEntity == nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "InteractionEvent with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToInteractionSession(interactionSessionEntity), nil
}

// InteractionSessionBySessionIdentifier is the resolver for the interactionSession_BySessionIdentifier field.
func (r *queryResolver) InteractionSessionBySessionIdentifier(ctx context.Context, sessionIdentifier string) (*model.InteractionSession, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.InteractionSessionBySessionIdentifier", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.sessionIdentifier", sessionIdentifier))

	interactionSessionEntity, err := r.Services.InteractionSessionService.GetInteractionSessionBySessionIdentifier(ctx, sessionIdentifier)
	if err != nil || interactionSessionEntity == nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "InteractionEvent with identifier %s not found", sessionIdentifier)
		return nil, err
	}
	return mapper.MapEntityToInteractionSession(interactionSessionEntity), nil
}

// InteractionEvent is the resolver for the interactionEvent field.
func (r *queryResolver) InteractionEvent(ctx context.Context, id string) (*model.InteractionEvent, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.InteractionEvent", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.interactionEventID", id))

	interactionEventEntity, err := r.Services.InteractionEventService.GetInteractionEventById(ctx, id)
	if err != nil || interactionEventEntity == nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "InteractionEvent with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToInteractionEvent(interactionEventEntity), nil
}

// InteractionEventByEventIdentifier is the resolver for the interactionEvent_ByEventIdentifier field.
func (r *queryResolver) InteractionEventByEventIdentifier(ctx context.Context, eventIdentifier string) (*model.InteractionEvent, error) {
	ctx, span := tracing.StartGraphQLTracerSpan(ctx, "QueryResolver.InteractionEventByEventIdentifier", graphql.GetOperationContext(ctx))
	defer span.Finish()
	tracing.SetDefaultResolverSpanTags(ctx, span)
	span.LogFields(log.String("request.eventIdentifier", eventIdentifier))

	interactionEventEntity, err := r.Services.InteractionEventService.GetInteractionEventByEventIdentifier(ctx, eventIdentifier)
	if err != nil || interactionEventEntity == nil {
		tracing.TraceErr(span, err)
		graphql.AddErrorf(ctx, "InteractionEvent with EventIdentifier %s not found", eventIdentifier)
		return nil, err
	}
	return mapper.MapEntityToInteractionEvent(interactionEventEntity), nil
}

// InteractionEvent returns generated.InteractionEventResolver implementation.
func (r *Resolver) InteractionEvent() generated.InteractionEventResolver {
	return &interactionEventResolver{r}
}

// InteractionSession returns generated.InteractionSessionResolver implementation.
func (r *Resolver) InteractionSession() generated.InteractionSessionResolver {
	return &interactionSessionResolver{r}
}

type interactionEventResolver struct{ *Resolver }
type interactionSessionResolver struct{ *Resolver }
