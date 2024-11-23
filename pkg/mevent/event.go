package mevent

import (
	"context"
	"fmt"
	"github.com/newrelic/go-agent/v3/newrelic"
	uuid "github.com/satori/go.uuid"
	"log/slog"
)

type SloggerEvent interface {
	ToSlogFields() []slog.Attr
}

type Event interface {
	Name() string
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
	logger   *slog.Logger
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
		if slogger, ok := event.(SloggerEvent); ok {
			e.logger.With(slog.Group(event.Name(), slogger.ToSlogFields())).Info("event published")
		} else {
			e.logger.With(slog.String("event.name", event.Name())).Info("event published")
		}
	} else {
		fmt.Println(fmt.Sprintf("event published: %s", event.Name()))
	}
	for _, handler := range e.handlers[event.Name()] {
		handler.Notify(event)
	}
}

type EventPublisherConfiguration func(e *Publisher)

func PublisherWithLogger(logger *slog.Logger) EventPublisherConfiguration {
	return func(e *Publisher) {
		e.logger = logger
	}
}
