package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/model"
)

// AttendedBy is the resolver for the attendedBy field.
func (r *meetingResolver) AttendedBy(ctx context.Context, obj *model.Meeting) ([]model.MeetingParticipant, error) {
	panic(fmt.Errorf("not implemented: AttendedBy - attendedBy"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *meetingResolver) CreatedBy(ctx context.Context, obj *model.Meeting) ([]model.MeetingParticipant, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// Includes is the resolver for the includes field.
func (r *meetingResolver) Includes(ctx context.Context, obj *model.Meeting) ([]*model.Attachment, error) {
	panic(fmt.Errorf("not implemented: Includes - includes"))
}

// DescribedBy is the resolver for the describedBy field.
func (r *meetingResolver) DescribedBy(ctx context.Context, obj *model.Meeting) ([]*model.Analysis, error) {
	panic(fmt.Errorf("not implemented: DescribedBy - describedBy"))
}

// Note is the resolver for the note field.
func (r *meetingResolver) Note(ctx context.Context, obj *model.Meeting) ([]*model.Note, error) {
	panic(fmt.Errorf("not implemented: Note - note"))
}

// Events is the resolver for the events field.
func (r *meetingResolver) Events(ctx context.Context, obj *model.Meeting) ([]*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: Events - events"))
}

// Recording is the resolver for the recording field.
func (r *meetingResolver) Recording(ctx context.Context, obj *model.Meeting) (*model.Attachment, error) {
	panic(fmt.Errorf("not implemented: Recording - recording"))
}

// MeetingCreate is the resolver for the meeting_Create field.
func (r *mutationResolver) MeetingCreate(ctx context.Context, meeting model.MeetingInput) (*model.Meeting, error) {
	panic(fmt.Errorf("not implemented: MeetingCreate - meeting_Create"))
}

// MeetingUpdate is the resolver for the meeting_Update field.
func (r *mutationResolver) MeetingUpdate(ctx context.Context, meetingID string, meeting model.MeetingUpdateInput) (*model.Meeting, error) {
	panic(fmt.Errorf("not implemented: MeetingUpdate - meeting_Update"))
}

// MeetingLinkAttendedBy is the resolver for the meeting_LinkAttendedBy field.
func (r *mutationResolver) MeetingLinkAttendedBy(ctx context.Context, meetingID string, participant model.MeetingParticipantInput) (*model.Meeting, error) {
	panic(fmt.Errorf("not implemented: MeetingLinkAttendedBy - meeting_LinkAttendedBy"))
}

// MeetingUnlinkAttendedBy is the resolver for the meeting_UnlinkAttendedBy field.
func (r *mutationResolver) MeetingUnlinkAttendedBy(ctx context.Context, meetingID string, participant model.MeetingParticipantInput) (*model.Meeting, error) {
	panic(fmt.Errorf("not implemented: MeetingUnlinkAttendedBy - meeting_UnlinkAttendedBy"))
}

// MeetingLinkAttachment is the resolver for the meeting_LinkAttachment field.
func (r *mutationResolver) MeetingLinkAttachment(ctx context.Context, meetingID string, attachmentID string) (*model.Meeting, error) {
	panic(fmt.Errorf("not implemented: MeetingLinkAttachment - meeting_LinkAttachment"))
}

// MeetingUnlinkAttachment is the resolver for the meeting_UnlinkAttachment field.
func (r *mutationResolver) MeetingUnlinkAttachment(ctx context.Context, meetingID string, attachmentID string) (*model.Meeting, error) {
	panic(fmt.Errorf("not implemented: MeetingUnlinkAttachment - meeting_UnlinkAttachment"))
}

// MeetingLinkRecording is the resolver for the meeting_LinkRecording field.
func (r *mutationResolver) MeetingLinkRecording(ctx context.Context, meetingID string, attachmentID string) (*model.Meeting, error) {
	panic(fmt.Errorf("not implemented: MeetingLinkRecording - meeting_LinkRecording"))
}

// MeetingUnlinkRecording is the resolver for the meeting_UnlinkRecording field.
func (r *mutationResolver) MeetingUnlinkRecording(ctx context.Context, meetingID string, attachmentID string) (*model.Meeting, error) {
	panic(fmt.Errorf("not implemented: MeetingUnlinkRecording - meeting_UnlinkRecording"))
}

// MeetingAddNewLocation is the resolver for the meeting_AddNewLocation field.
func (r *mutationResolver) MeetingAddNewLocation(ctx context.Context, meetingID string) (*model.Location, error) {
	panic(fmt.Errorf("not implemented: MeetingAddNewLocation - meeting_AddNewLocation"))
}

// Meeting is the resolver for the meeting field.
func (r *queryResolver) Meeting(ctx context.Context, id string) (*model.Meeting, error) {
	panic(fmt.Errorf("not implemented: Meeting - meeting"))
}

// Meeting returns generated.MeetingResolver implementation.
func (r *Resolver) Meeting() generated.MeetingResolver { return &meetingResolver{r} }

type meetingResolver struct{ *Resolver }
