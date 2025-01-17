package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/model"
)

// Tags is the resolver for the tags field.
func (r *issueResolver) Tags(ctx context.Context, obj *model.Issue) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented: Tags - tags"))
}

// MentionedByNotes is the resolver for the mentionedByNotes field.
func (r *issueResolver) MentionedByNotes(ctx context.Context, obj *model.Issue) ([]*model.Note, error) {
	panic(fmt.Errorf("not implemented: MentionedByNotes - mentionedByNotes"))
}

// InteractionEvents is the resolver for the interactionEvents field.
func (r *issueResolver) InteractionEvents(ctx context.Context, obj *model.Issue) ([]*model.InteractionEvent, error) {
	panic(fmt.Errorf("not implemented: InteractionEvents - interactionEvents"))
}

// ExternalLinks is the resolver for the externalLinks field.
func (r *issueResolver) ExternalLinks(ctx context.Context, obj *model.Issue) ([]*model.ExternalSystem, error) {
	panic(fmt.Errorf("not implemented: ExternalLinks - externalLinks"))
}

// SubmittedBy is the resolver for the submittedBy field.
func (r *issueResolver) SubmittedBy(ctx context.Context, obj *model.Issue) (model.IssueParticipant, error) {
	panic(fmt.Errorf("not implemented: SubmittedBy - submittedBy"))
}

// ReportedBy is the resolver for the reportedBy field.
func (r *issueResolver) ReportedBy(ctx context.Context, obj *model.Issue) (model.IssueParticipant, error) {
	panic(fmt.Errorf("not implemented: ReportedBy - reportedBy"))
}

// AssignedTo is the resolver for the assignedTo field.
func (r *issueResolver) AssignedTo(ctx context.Context, obj *model.Issue) ([]model.IssueParticipant, error) {
	panic(fmt.Errorf("not implemented: AssignedTo - assignedTo"))
}

// FollowedBy is the resolver for the followedBy field.
func (r *issueResolver) FollowedBy(ctx context.Context, obj *model.Issue) ([]model.IssueParticipant, error) {
	panic(fmt.Errorf("not implemented: FollowedBy - followedBy"))
}

// Issue is the resolver for the issue field.
func (r *queryResolver) Issue(ctx context.Context, id string) (*model.Issue, error) {
	panic(fmt.Errorf("not implemented: Issue - issue"))
}

// Issue returns generated.IssueResolver implementation.
func (r *Resolver) Issue() generated.IssueResolver { return &issueResolver{r} }

type issueResolver struct{ *Resolver }
