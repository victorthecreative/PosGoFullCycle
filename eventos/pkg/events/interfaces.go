package events

import "time"

type EventInterface interface {
	getName() string
	getDataTime() time.Time
	getPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface)
}

type EventDispatcherInterface interface {
	// Register an event with a handler
	Register(eventName string, handler EventHandlerInterface) error
	// Dispatch an event
	Dispatch(event EventInterface) error
	// Remove an event handler
	Remove(eventName string, handler EventHandlerInterface) error
	// Check if an event has a handler
	Has(eventName string, handler EventHandlerInterface) bool
	// Get all handlers for an event
	Claer() error
}
