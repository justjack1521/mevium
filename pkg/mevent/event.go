package mevent

import (
	"context"
	"github.com/newrelic/go-agent/v3/newrelic"
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

type ContextEvent interface {
	Event
	Context() context.Context
}

type NewRelicEvent interface {
	Event
	Transaction() *newrelic.Transaction
}

type Handler interface {
	Notify(event Event)
}

type Publisher struct {
	handlers map[string][]Handler
	logger   *logrus.Logger
}

func NewPublisher(options ...EventPublisherConfiguration) *Publisher {
	publisher := &Publisher{handlers: map[string][]Handler{}}
	for _, opt := range options {
		opt(publisher)
	}
	return publisher
}

func (e *Publisher) Subscribe(handler Handler, events ...Event) {
	for _, event := range events {
		handlers := e.handlers[event.Name()]
		handlers = append(handlers, handler)
		e.handlers[event.Name()] = handlers
	}
}

func (e *Publisher) Notify(event Event) {
	if e.logger != nil {
		e.logger.WithFields(event.ToLogFields()).Info("Event Published")
	}
	for _, handler := range e.handlers[event.Name()] {
		handler.Notify(event)
	}
}

type EventPublisherConfiguration func(e *Publisher)

func PublisherWithLogger(logger *logrus.Logger) EventPublisherConfiguration {
	return func(e *Publisher) {
		e.logger = logger
	}
}
