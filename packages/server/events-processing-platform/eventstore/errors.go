package eventstore

import (
	"fmt"
	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/pkg/errors"
)

var (
	ErrInvalidEventType    = InvalidEventTypeError{}
	ErrAlreadyExists       = errors.New("already exists")
	ErrAggregateNotFound   = errors.New("aggregate not found")
	ErrInvalidCommandType  = errors.New("invalid command type")
	ErrInvalidAggregate    = errors.New("invalid aggregate")
	ErrInvalidAggregateID  = errors.New("invalid aggregate id")
	ErrInvalidEventVersion = errors.New("invalid event version")
	ErrMissingTenant       = errors.New("missing tenant")
)

type InvalidEventTypeError struct {
	EventType string
}

func (e InvalidEventTypeError) Error() string {
	return fmt.Sprintf("invalid event type: %s", e.EventType)
}

func IsEventStoreErrorCodeResourceNotFound(err error) bool {
	esdbErr, ok := esdb.FromError(err)
	if ok {
		return false
	}
	errorCode := esdbErr.Code()
	if errorCode == esdb.ErrorCodeResourceNotFound {
		return true
	}
	return false
}

func IsEventStoreErrorCodeResourceAlreadyExists(err error) bool {
	esdbErr, ok := esdb.FromError(err)
	if ok {
		return false
	}
	errorCode := esdbErr.Code()
	if errorCode == esdb.ErrorCodeResourceAlreadyExists {
		return true
	}
	return false
}
