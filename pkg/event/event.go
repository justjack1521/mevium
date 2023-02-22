package event

type Event interface {
	Name() string
}

type Handler interface {
	Notify(event Event)
}

type Publisher struct {
	handlers map[string][]Handler
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
