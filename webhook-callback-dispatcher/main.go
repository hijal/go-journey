package main

import (
	"fmt"
	"sync"
)

type Event struct {
	Name    string
	Payload map[string]any
}

type EventBus struct {
	mu        sync.RWMutex
	listeners map[string][]func(Event)
}

func NewEventBus() *EventBus {
	return &EventBus{
		listeners: make(map[string][]func(Event)),
	}
}

func (eb *EventBus) Subscribe(eventName string, callback func(Event)) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.listeners[eventName] = append(eb.listeners[eventName], callback)
}

func (eb *EventBus) Publish(event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	if callbacks, ok := eb.listeners[event.Name]; ok {
		for _, cb := range callbacks {
			cb(event)
		}
	}
}

func main() {
	bus := NewEventBus()

	bus.Subscribe("user_register", func(e Event) {
		fmt.Printf("[Webhook 1] sending welcome email to %v\n", e.Payload["email"])
	})

	bus.Subscribe("user_register", func(e Event) {
		fmt.Printf("[Webhook 2] Notifying slack channel about user %v\n", e.Payload["username"])
	})

	bus.Publish(Event{
		Name: "user_register",
		Payload: map[string]any{
			"username": "johndoe",
			"email":    "john@example.com",
		},
	})
}
