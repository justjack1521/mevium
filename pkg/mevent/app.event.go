package mevent

import "github.com/sirupsen/logrus"

type ApplicationEvent interface {
	Event
}

type ApplicationStartEvent struct {
}

func (e ApplicationStartEvent) Name() string {
	return "event.application.start"
}

func (e ApplicationStartEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{"event_name": e.Name()}
}

type ApplicationShutdownEvent struct {
}

func (e ApplicationShutdownEvent) Name() string {
	return "event.application.shutdown"
}

func (e ApplicationShutdownEvent) ToLogFields() logrus.Fields {
	return logrus.Fields{"event_name": e.Name()}
}
