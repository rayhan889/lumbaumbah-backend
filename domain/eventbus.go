package domain

import "fmt"

type Event interface {
	Eventname() string
}

type EventHandler func(Event)

var eventHanlders = map[string][]EventHandler{}

func RegisterHandler(eventName string, handler EventHandler) {
	eventHanlders[eventName] = append(eventHanlders[eventName], handler)
}

func Emit(event Event) {
	fmt.Printf("Emitting event: %s\n", event.Eventname())
	fmt.Printf("Event data: %+v\n", event)
	if handlers, ok := eventHanlders[event.Eventname()]; ok {
		for _, h := range handlers {
			go h(event)
		}
	}
}
