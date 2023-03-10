package mevent

import (
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Event interface {
	Name() string
	ToLogFields() logrus.Fields
}

type ClientEvent interface {
	Event
	ClientID() uuid.UUID
}

type PlayerEvent interface {
	Event
	PlayerID() uuid.UUID
}

type Handler interface {
	Notify(event Event)
}

type Publisher struct {
	handlers map[string][]Handler
}

func NewPublisher() *Publisher {
	return &Publisher{handlers: map[string][]Handler{}}
}

func (e *Publisher) Subscribe(handler Handler, events ...Event) {
	for _, event := range events {
		handlers := e.handlers[event.Name()]
		handlers = append(handlers, handler)
		e.handlers[event.Name()] = handlers
	}
}

func (e *Publisher) Notify(event Event) {
	for _, handler := range e.handlers[event.Name()] {
		handler.Notify(event)
	}
}
