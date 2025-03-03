package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

// EventDispatcher is responsible for managing event handlers and dispatching events to them.
// It maintains a map where the keys are event names and the values are slices of EventHandlerInterface.
// This allows multiple handlers to be registered for a single event.
type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

// NewEventDispatcher creates a new EventDispatcher instance.
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Register adds a handler for a given event.
func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {

	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}
