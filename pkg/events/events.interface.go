package events

import (
	"context"
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup, ctx context.Context)
}

type EventDispatcherInterface interface {
	Register(name string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(name string) error
	Has(name string) bool
	Clear()
}
