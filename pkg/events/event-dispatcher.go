package events

import (
	"context"
	"errors"
	"sync"
)

var errHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string]EventHandlerInterface),
	}
}

func (e *EventDispatcher) Register(name string, handler EventHandlerInterface) error {
	if _, ok := e.handlers[name]; ok {
		return errHandlerAlreadyRegistered
	}

	e.handlers[name] = handler

	return nil
}

func (e *EventDispatcher) Dispatch(event EventInterface) error {
	if _, ok := e.handlers[event.GetName()]; !ok {
		return nil
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ctx := context.Background()

	e.handlers[event.GetName()].Handle(event, wg, ctx)

	wg.Wait()

	return nil
}

func (e *EventDispatcher) Remove(name string) error {
	if _, ok := e.handlers[name]; !ok {
		return nil
	}

	delete(e.handlers, name)

	return nil
}

func (e *EventDispatcher) Has(name string) bool {
	if _, ok := e.handlers[name]; !ok {
		return false
	}

	return true
}

func (e *EventDispatcher) Clear() {
	e.handlers = make(map[string]EventHandlerInterface)
}
