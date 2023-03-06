package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

// Tags is the resolver for the tags field.
func (r *ticketResolver) Tags(ctx context.Context, obj *model.Ticket) ([]*model.Tag, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	tagEntities, err := dataloader.For(ctx).GetTagsForTickets(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get tags for contact %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToTags(tagEntities), nil
}

// Notes is the resolver for the notes field.
func (r *ticketResolver) Notes(ctx context.Context, obj *model.Ticket) ([]*model.Note, error) {
	defer func(start time.Time) {
		utils.LogMethodExecution(start, utils.GetFunctionName())
	}(time.Now())

	noteEntities, err := dataloader.For(ctx).GetNotesForTicket(ctx, obj.ID)
	if err != nil {
		graphql.AddErrorf(ctx, "Failed to get notes for ticket %s", obj.ID)
		return nil, err
	}
	return mapper.MapEntitiesToNotes(noteEntities), nil
}

// Ticket returns generated.TicketResolver implementation.
func (r *Resolver) Ticket() generated.TicketResolver { return &ticketResolver{r} }

type ticketResolver struct{ *Resolver }