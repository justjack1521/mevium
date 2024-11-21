package mevent

type ApplicationEvent interface {
	Event
}

type ApplicationStartEvent struct {
}

func (e ApplicationStartEvent) Name() string {
	return "event.application.start"
}

type ApplicationShutdownEvent struct {
}

func (e ApplicationShutdownEvent) Name() string {
	return "event.application.shutdown"
}
