package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// AttendedBy is the resolver for the attendedBy field.
func (r *meetingResolver) AttendedBy(ctx context.Context, obj *model.Meeting) ([]model.MeetingParticipant, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	participantEntities, err := dataloader.For(ctx).GetAttendedByParticipantsForMeeting(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get participants for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToMeetingParticipants(participantEntities), nil
}

// CreatedBy is the resolver for the createdBy field.
func (r *meetingResolver) CreatedBy(ctx context.Context, obj *model.Meeting) ([]model.MeetingParticipant, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	participantEntities, err := dataloader.For(ctx).GetCreatedByParticipantsForMeeting(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get participants for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToMeetingParticipants(participantEntities), nil
}

// Includes is the resolver for the includes field.
func (r *meetingResolver) Includes(ctx context.Context, obj *model.Meeting) ([]*model.Attachment, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())
	entities, err := dataloader.For(ctx).GetAttachmentsForMeeting(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get attachment entities for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToAttachment(entities), nil
}

// DescribedBy is the resolver for the describedBy field.
func (r *meetingResolver) DescribedBy(ctx context.Context, obj *model.Meeting) ([]*model.Analysis, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	analysisEntities, err := dataloader.For(ctx).GetDescribedByForMeeting(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get analysis for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToAnalysis(analysisEntities), nil
}

// Note is the resolver for the note field.
func (r *meetingResolver) Note(ctx context.Context, obj *model.Meeting) ([]*model.Note, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	notesForMeeting, err := dataloader.For(ctx).GetNotesForMeeting(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get notes for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToNotes(notesForMeeting), nil
}

// Events is the resolver for the events field.
func (r *meetingResolver) Events(ctx context.Context, obj *model.Meeting) ([]*model.InteractionEvent, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	interactionEventEntities, err := dataloader.For(ctx).GetInteractionEventsForMeeting(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get interaction events for meeting %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToInteractionEvents(interactionEventEntities), nil
}

// Recording is the resolver for the recording field.
func (r *meetingResolver) Recording(ctx context.Context, obj *model.Meeting) (*model.Attachment, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())
	recording := repository.INCLUDE_NATURE_RECORDING
	entities, err := r.Services.AttachmentService.GetAttachmentsForNode(ctx, repository.INCLUDED_BY_MEETING, &recording, []string{obj.ID})
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get attachment entities for meeting %s", obj.ID)
		return nil, err
	}
	attachment := mapper.MapEntitiesToAttachment(entities)

	if len(attachment) == 0 {
		return nil, nil
	}

	return attachment[0], nil
}

// MeetingCreate is the resolver for the meeting_Create field.
func (r *mutationResolver) MeetingCreate(ctx context.Context, meeting model.MeetingInput) (*model.Meeting, error) {
	meetingEntity, err := r.Services.MeetingService.Create(ctx,
		&service.MeetingCreateData{
			MeetingEntity: mapper.MapMeetingToEntity(&meeting),
			CreatedBy:     service.MapMeetingParticipantInputListToParticipant(meeting.CreatedBy),
			AttendedBy:    service.MapMeetingParticipantInputListToParticipant(meeting.AttendedBy),
			NoteInput:     meeting.Note,
		})
	if err != nil {
		graphql.AddErrorf(ctx, "failed to create meeting")
		return nil, err
	}
	newMeeting := mapper.MapEntityToMeeting(meetingEntity)
	return newMeeting, nil
}

// MeetingUpdate is the resolver for the meeting_Update field.
func (r *mutationResolver) MeetingUpdate(ctx context.Context, meetingID string, meeting model.MeetingUpdateInput) (*model.Meeting, error) {
	input := &service.MeetingUpdateData{
		MeetingEntity: mapper.MapMeetingInputToEntity(&meeting),
		NoteEntity:    mapper.MapNoteUpdateInputToEntity(meeting.Note),
	}
	input.MeetingEntity.Id = meetingID
	meetingEntity, err := r.Services.MeetingService.Update(ctx, input)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to update meeting")
		return nil, err
	}
	interactionEvent := mapper.MapEntityToMeeting(meetingEntity)
	return interactionEvent, nil
}

// MeetingLinkAttendedBy is the resolver for the meeting_LinkAttendedBy field.
func (r *mutationResolver) MeetingLinkAttendedBy(ctx context.Context, meetingID string, participant model.MeetingParticipantInput) (*model.Meeting, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	err := r.Services.MeetingService.LinkAttendedBy(ctx, meetingID, service.MapMeetingParticipantInputToParticipant(&participant))
	if err != nil {
		return nil, err
	}

	meeting, err := r.Services.MeetingService.GetMeetingById(ctx, meetingID)
	if err != nil {
		return nil, err
	}

	return mapper.MapEntityToMeeting(meeting), nil
}

// MeetingUnlinkAttendedBy is the resolver for the meeting_UnlinkAttendedBy field.
func (r *mutationResolver) MeetingUnlinkAttendedBy(ctx context.Context, meetingID string, participant model.MeetingParticipantInput) (*model.Meeting, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	err := r.Services.MeetingService.UnlinkAttendedBy(ctx, meetingID, service.MapMeetingParticipantInputToParticipant(&participant))
	if err != nil {
		return nil, err
	}

	meeting, err := r.Services.MeetingService.GetMeetingById(ctx, meetingID)
	if err != nil {
		return nil, err
	}

	return mapper.MapEntityToMeeting(meeting), nil
}

// MeetingLinkAttachment is the resolver for the meeting_LinkAttachment field.
func (r *mutationResolver) MeetingLinkAttachment(ctx context.Context, meetingID string, attachmentID string) (*model.Meeting, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())
	meeting, err := r.Services.MeetingService.LinkAttachment(ctx, meetingID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToMeeting(meeting), nil
}

// MeetingUnlinkAttachment is the resolver for the meeting_UnlinkAttachment field.
func (r *mutationResolver) MeetingUnlinkAttachment(ctx context.Context, meetingID string, attachmentID string) (*model.Meeting, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())
	meeting, err := r.Services.MeetingService.UnlinkAttachment(ctx, meetingID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToMeeting(meeting), nil
}

// MeetingLinkRecording is the resolver for the meeting_LinkRecording field.
func (r *mutationResolver) MeetingLinkRecording(ctx context.Context, meetingID string, attachmentID string) (*model.Meeting, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())
	meeting, err := r.Services.MeetingService.LinkRecordingAttachment(ctx, meetingID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToMeeting(meeting), nil
}

// MeetingUnlinkRecording is the resolver for the meeting_UnlinkRecording field.
func (r *mutationResolver) MeetingUnlinkRecording(ctx context.Context, meetingID string, attachmentID string) (*model.Meeting, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())
	meeting, err := r.Services.MeetingService.UnlinkRecordingAttachment(ctx, meetingID, attachmentID)
	if err != nil {
		return nil, err
	}
	return mapper.MapEntityToMeeting(meeting), nil
}

// Meeting is the resolver for the meeting field.
func (r *queryResolver) Meeting(ctx context.Context, id string) (*model.Meeting, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	meetingEntity, err := r.Services.MeetingService.GetMeetingById(ctx, id)
	if err != nil || meetingEntity == nil {
		graphql.AddErrorf(ctx, "Meeting with id %s not found", id)
		return nil, err
	}
	return mapper.MapEntityToMeeting(meetingEntity), nil
}

// Meeting returns generated.MeetingResolver implementation.
func (r *Resolver) Meeting() generated.MeetingResolver { return &meetingResolver{r} }

type meetingResolver struct{ *Resolver }