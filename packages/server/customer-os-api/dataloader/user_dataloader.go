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

func (i *Loaders) GetUsersForEmail(ctx context.Context, emailID string) (*entity.UserEntities, error) {
	thunk := i.UsersForEmail.Load(ctx, dataloader.StringKey(emailID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	resultObj := result.(entity.UserEntities)
	return &resultObj, nil
}

func (i *Loaders) GetUsersForPhoneNumber(ctx context.Context, phoneNumberID string) (*entity.UserEntities, error) {
	thunk := i.UsersForPhoneNumber.Load(ctx, dataloader.StringKey(phoneNumberID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	resultObj := result.(entity.UserEntities)
	return &resultObj, nil
}

func (i *Loaders) GetUsersForPlayer(ctx context.Context, playerID string) (*entity.UserEntities, error) {
	thunk := i.UsersForPlayer.Load(ctx, dataloader.StringKey(playerID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	resultObj := result.(entity.UserEntities)
	return &resultObj, nil
}

func (i *Loaders) GetUserOwnerForOrganization(ctx context.Context, organizationID string) (*entity.UserEntity, error) {
	thunk := i.UserOwnerForOrganization.Load(ctx, dataloader.StringKey(organizationID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	return result.(*entity.UserEntity), nil
}

func (i *Loaders) GetUser(ctx context.Context, userId string) (*entity.UserEntity, error) {
	thunk := i.User.Load(ctx, dataloader.StringKey(userId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	return result.(*entity.UserEntity), nil
}

func (i *Loaders) GetUserAuthorForLogEntry(ctx context.Context, logEntryId string) (*entity.UserEntity, error) {
	thunk := i.UserAuthorForLogEntry.Load(ctx, dataloader.StringKey(logEntryId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	return result.(*entity.UserEntity), nil
}

func (b *userBatcher) getUsersForEmails(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserDataLoader.getUsersForEmails", opentracing.ChildOf(tracing.ExtractSpanCtx(ctx)))
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("keys", keys), log.Int("keys_length", len(keys)))

	ids, keyOrder := sortKeys(keys)

	userEntitiesPtr, err := b.userService.GetUsersForEmails(ctx, ids)
	if err != nil {
		tracing.TraceErr(span, err)
		// check if context deadline exceeded error occurred
		if ctx.Err() == context.DeadlineExceeded {
			return []*dataloader.Result{{Data: nil, Error: errors.New("deadline exceeded to get users for emails")}}
		}
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	userEntitiesByEmailId := make(map[string]entity.UserEntities)
	for _, val := range *userEntitiesPtr {
		if list, ok := userEntitiesByEmailId[val.DataloaderKey]; ok {
			userEntitiesByEmailId[val.DataloaderKey] = append(list, val)
		} else {
			userEntitiesByEmailId[val.DataloaderKey] = entity.UserEntities{val}
		}
	}

	// construct an output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	for emailId, record := range userEntitiesByEmailId {
		if ix, ok := keyOrder[emailId]; ok {
			results[ix] = &dataloader.Result{Data: record, Error: nil}
			delete(keyOrder, emailId)
		}
	}
	for _, ix := range keyOrder {
		results[ix] = &dataloader.Result{Data: entity.UserEntities{}, Error: nil}
	}

	if err = assertEntitiesType(results, reflect.TypeOf(entity.UserEntities{})); err != nil {
		tracing.TraceErr(span, err)
		return []*dataloader.Result{{nil, err}}
	}

	span.LogFields(log.Object("output - results_length", len(results)))

	return results
}

func (b *userBatcher) getUsersForPhoneNumbers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserDataLoader.getUsersForPhoneNumbers", opentracing.ChildOf(tracing.ExtractSpanCtx(ctx)))
	defer span.Finish()
	span.LogFields(log.Object("keys", keys), log.Int("keys_length", len(keys)))

	ids, keyOrder := sortKeys(keys)

	userEntitiesPtr, err := b.userService.GetUsersForPhoneNumbers(ctx, ids)
	if err != nil {
		tracing.TraceErr(span, err)
		// check if context deadline exceeded error occurred
		if ctx.Err() == context.DeadlineExceeded {
			return []*dataloader.Result{{Data: nil, Error: errors.New("deadline exceeded to get users for phone numbers")}}
		}
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	userEntitiesByPhoneNumberId := make(map[string]entity.UserEntities)
	for _, val := range *userEntitiesPtr {
		if list, ok := userEntitiesByPhoneNumberId[val.DataloaderKey]; ok {
			userEntitiesByPhoneNumberId[val.DataloaderKey] = append(list, val)
		} else {
			userEntitiesByPhoneNumberId[val.DataloaderKey] = entity.UserEntities{val}
		}
	}

	// construct an output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	for phoneNumberId, record := range userEntitiesByPhoneNumberId {
		if ix, ok := keyOrder[phoneNumberId]; ok {
			results[ix] = &dataloader.Result{Data: record, Error: nil}
			delete(keyOrder, phoneNumberId)
		}
	}
	for _, ix := range keyOrder {
		results[ix] = &dataloader.Result{Data: entity.UserEntities{}, Error: nil}
	}

	if err = assertEntitiesType(results, reflect.TypeOf(entity.UserEntities{})); err != nil {
		tracing.TraceErr(span, err)
		return []*dataloader.Result{{nil, err}}
	}

	span.LogFields(log.Object("output - results_length", len(results)))

	return results
}

func (b *userBatcher) getUsersForPlayers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserDataLoader.getUsersForPlayers", opentracing.ChildOf(tracing.ExtractSpanCtx(ctx)))
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("keys", keys), log.Int("keys_length", len(keys)))

	ids, keyOrder := sortKeys(keys)

	userEntitiesPtr, err := b.userService.GetUsersForPlayers(ctx, ids)
	if err != nil {
		tracing.TraceErr(span, err)
		// check if context deadline exceeded error occurred
		if ctx.Err() == context.DeadlineExceeded {
			return []*dataloader.Result{{Data: nil, Error: errors.New("deadline exceeded to get users for players")}}
		}
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	userEntitiesByPlayerId := make(map[string]entity.UserEntities)
	for _, val := range *userEntitiesPtr {
		if list, ok := userEntitiesByPlayerId[val.DataloaderKey]; ok {
			userEntitiesByPlayerId[val.DataloaderKey] = append(list, val)
		} else {
			userEntitiesByPlayerId[val.DataloaderKey] = entity.UserEntities{val}
		}
	}

	// construct an output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	for phoneNumberId, record := range userEntitiesByPlayerId {
		if ix, ok := keyOrder[phoneNumberId]; ok {
			results[ix] = &dataloader.Result{Data: record, Error: nil}
			delete(keyOrder, phoneNumberId)
		}
	}
	for _, ix := range keyOrder {
		results[ix] = &dataloader.Result{Data: entity.UserEntities{}, Error: nil}
	}

	if err = assertEntitiesType(results, reflect.TypeOf(entity.UserEntities{})); err != nil {
		tracing.TraceErr(span, err)
		return []*dataloader.Result{{nil, err}}
	}

	span.LogFields(log.Object("output - results_length", len(results)))

	return results
}

func (b *userBatcher) getUserOwnersForOrganizations(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserDataLoader.getUserOwnersForOrganizations", opentracing.ChildOf(tracing.ExtractSpanCtx(ctx)))
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("keys", keys), log.Int("keys_length", len(keys)))

	ids, keyOrder := sortKeys(keys)

	userEntities, err := b.userService.GetUserOwnersForOrganizations(ctx, ids)
	if err != nil {
		tracing.TraceErr(span, err)
		// check if context deadline exceeded error occurred
		if ctx.Err() == context.DeadlineExceeded {
			return []*dataloader.Result{{Data: nil, Error: errors.New("deadline exceeded to get users for organizations")}}
		}
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	userEntityByOrganizationId := make(map[string]entity.UserEntity)
	for _, val := range *userEntities {
		userEntityByOrganizationId[val.DataloaderKey] = val
	}

	// construct an output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	for organizationID, _ := range userEntityByOrganizationId {
		if ix, ok := keyOrder[organizationID]; ok {
			val := userEntityByOrganizationId[organizationID]
			results[ix] = &dataloader.Result{Data: &val, Error: nil}
			delete(keyOrder, organizationID)
		}
	}
	for _, ix := range keyOrder {
		results[ix] = &dataloader.Result{Data: nil, Error: nil}
	}

	if err = assertEntitiesPtrType(results, reflect.TypeOf(entity.UserEntity{}), true); err != nil {
		tracing.TraceErr(span, err)
		return []*dataloader.Result{{nil, err}}
	}

	span.LogFields(log.Object("output - results_length", len(results)))

	return results
}

func (b *userBatcher) getUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserDataLoader.getUsers", opentracing.ChildOf(tracing.ExtractSpanCtx(ctx)))
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("keys", keys), log.Int("keys_length", len(keys)))

	ids, keyOrder := sortKeys(keys)

	userEntities, err := b.userService.GetUsers(ctx, ids)
	if err != nil {
		tracing.TraceErr(span, err)
		// check if context deadline exceeded error occurred
		if ctx.Err() == context.DeadlineExceeded {
			return []*dataloader.Result{{Data: nil, Error: errors.New("deadline exceeded to get users")}}
		}
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	userEntityById := make(map[string]entity.UserEntity)
	for _, val := range *userEntities {
		userEntityById[val.Id] = val
	}

	// construct an output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	for id, _ := range userEntityById {
		if ix, ok := keyOrder[id]; ok {
			val := userEntityById[id]
			results[ix] = &dataloader.Result{Data: &val, Error: nil}
			delete(keyOrder, id)
		}
	}
	for _, ix := range keyOrder {
		results[ix] = &dataloader.Result{Data: nil, Error: nil}
	}

	if err = assertEntitiesPtrType(results, reflect.TypeOf(entity.UserEntity{}), true); err != nil {
		tracing.TraceErr(span, err)
		return []*dataloader.Result{{nil, err}}
	}

	span.LogFields(log.Object("output - results_length", len(results)))

	return results
}

func (b *userBatcher) getUserAuthorsForLogEntries(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserDataLoader.getUserAuthorsForLogEntries", opentracing.ChildOf(tracing.ExtractSpanCtx(ctx)))
	defer span.Finish()
	tracing.SetDefaultServiceSpanTags(ctx, span)
	span.LogFields(log.Object("keys", keys), log.Int("keys_length", len(keys)))

	ids, keyOrder := sortKeys(keys)

	userEntities, err := b.userService.GetUserAuthorsForLogEntries(ctx, ids)
	if err != nil {
		tracing.TraceErr(span, err)
		// check if context deadline exceeded error occurred
		if ctx.Err() == context.DeadlineExceeded {
			return []*dataloader.Result{{Data: nil, Error: errors.New("deadline exceeded to get users for log entries")}}
		}
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	userEntityByLogEntryId := make(map[string]entity.UserEntity)
	for _, val := range *userEntities {
		userEntityByLogEntryId[val.DataloaderKey] = val
	}

	// construct an output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	for logEntryId, _ := range userEntityByLogEntryId {
		if ix, ok := keyOrder[logEntryId]; ok {
			val := userEntityByLogEntryId[logEntryId]
			results[ix] = &dataloader.Result{Data: &val, Error: nil}
			delete(keyOrder, logEntryId)
		}
	}
	for _, ix := range keyOrder {
		results[ix] = &dataloader.Result{Data: nil, Error: nil}
	}

	if err = assertEntitiesPtrType(results, reflect.TypeOf(entity.UserEntity{}), true); err != nil {
		tracing.TraceErr(span, err)
		return []*dataloader.Result{{nil, err}}
	}

	span.LogFields(log.Object("output - results_length", len(results)))

	return results
}
