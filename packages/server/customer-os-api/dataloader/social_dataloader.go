package dataloader

import (
	"context"
	"errors"
	"github.com/graph-gophers/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"reflect"
)

func (i *Loaders) GetSocialsForContact(ctx context.Context, contactId string) (*entity.SocialEntities, error) {
	thunk := i.SocialsForContact.Load(ctx, dataloader.StringKey(contactId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	resultObj := result.(entity.SocialEntities)
	return &resultObj, nil
}

func (i *Loaders) GetSocialsForOrganization(ctx context.Context, organizationId string) (*entity.SocialEntities, error) {
	thunk := i.SocialsForOrganization.Load(ctx, dataloader.StringKey(organizationId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	resultObj := result.(entity.SocialEntities)
	return &resultObj, nil
}

func (b *socialBatcher) getSocialsForContacts(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SocialDataLoader.getSocialsForContacts", opentracing.ChildOf(tracing.ExtractSpanCtx(ctx)))
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("keys", keys), log.Int("keys_length", len(keys)))

	ids, keyOrder := sortKeys(keys)

	socialEntitiesPtr, err := b.socialService.GetAllForEntities(ctx, entity.CONTACT, ids)
	if err != nil {
		tracing.TraceErr(span, err)
		// check if context deadline exceeded error occurred
		if ctx.Err() == context.DeadlineExceeded {
			return []*dataloader.Result{{Data: nil, Error: errors.New("deadline exceeded to get socials for contacts")}}
		}
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	socialEntitiesGrouped := make(map[string]entity.SocialEntities)
	for _, val := range *socialEntitiesPtr {
		if list, ok := socialEntitiesGrouped[val.DataloaderKey]; ok {
			socialEntitiesGrouped[val.DataloaderKey] = append(list, val)
		} else {
			socialEntitiesGrouped[val.DataloaderKey] = entity.SocialEntities{val}
		}
	}

	// construct an output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	for contactId, record := range socialEntitiesGrouped {
		ix, ok := keyOrder[contactId]
		if ok {
			results[ix] = &dataloader.Result{Data: record, Error: nil}
			delete(keyOrder, contactId)
		}
	}
	for _, ix := range keyOrder {
		results[ix] = &dataloader.Result{Data: entity.SocialEntities{}, Error: nil}
	}

	if err = assertEntitiesType(results, reflect.TypeOf(entity.SocialEntities{})); err != nil {
		tracing.TraceErr(span, err)
		return []*dataloader.Result{{nil, err}}
	}

	span.LogFields(log.Int("results_length", len(results)))

	return results
}

func (b *socialBatcher) getSocialsForOrganizations(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SocialDataLoader.getSocialsForOrganizations", opentracing.ChildOf(tracing.ExtractSpanCtx(ctx)))
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("keys", keys), log.Int("keys_length", len(keys)))

	ids, keyOrder := sortKeys(keys)

	socialEntitiesPtr, err := b.socialService.GetAllForEntities(ctx, entity.ORGANIZATION, ids)
	if err != nil {
		tracing.TraceErr(span, err)
		// check if context deadline exceeded error occurred
		if ctx.Err() == context.DeadlineExceeded {
			return []*dataloader.Result{{Data: nil, Error: errors.New("deadline exceeded to get socials for organizations")}}
		}
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	socialEntitiesGrouped := make(map[string]entity.SocialEntities)
	for _, val := range *socialEntitiesPtr {
		if list, ok := socialEntitiesGrouped[val.DataloaderKey]; ok {
			socialEntitiesGrouped[val.DataloaderKey] = append(list, val)
		} else {
			socialEntitiesGrouped[val.DataloaderKey] = entity.SocialEntities{val}
		}
	}

	// construct an output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	for organizationId, record := range socialEntitiesGrouped {
		ix, ok := keyOrder[organizationId]
		if ok {
			results[ix] = &dataloader.Result{Data: record, Error: nil}
			delete(keyOrder, organizationId)
		}
	}
	for _, ix := range keyOrder {
		results[ix] = &dataloader.Result{Data: entity.SocialEntities{}, Error: nil}
	}

	if err = assertEntitiesType(results, reflect.TypeOf(entity.SocialEntities{})); err != nil {
		tracing.TraceErr(span, err)
		return []*dataloader.Result{{nil, err}}
	}

	span.LogFields(log.Int("results_length", len(results)))

	return results
}
