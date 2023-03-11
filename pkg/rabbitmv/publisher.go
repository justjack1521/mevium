package rabbitmv

import (
	"context"
	"github.com/newrelic/go-agent/v3/newrelic"
	uuid "github.com/satori/go.uuid"
	"github.com/wagslane/go-rabbitmq"
)

type StandardPublisher struct {
	exchange Exchange
	kind     ExchangeKind
	closed   bool
	actual   *rabbitmq.Publisher
}

func (s *StandardPublisher) Close() {
	if s.closed == true {
		return
	}
	s.closed = true
	s.actual.Close()
}

func (s *StandardPublisher) Publish(bytes []byte, client uuid.UUID, key RoutingKey, options ...func(*rabbitmq.PublishOptions)) error {
	options = append(options, rabbitmq.WithPublishOptionsExchange(string(s.exchange)), rabbitmq.WithPublishOptionsHeaders(ClientPublishingTable(client)))
	return s.actual.Publish(
		bytes,
		[]string{string(key)},
		options...,
	)
}

func (s *StandardPublisher) PublishWithContext(ctx context.Context, bytes []byte, client uuid.UUID, key RoutingKey) error {
	txn := newrelic.FromContext(ctx)
	return s.PublishWithSegment(txn, bytes, client, key)
}

func (s *StandardPublisher) PublishWithSegment(txn *newrelic.Transaction, bytes []byte, client uuid.UUID, key RoutingKey) error {
	segment := newrelic.MessageProducerSegment{
		StartTime:            txn.StartSegmentNow(),
		Library:              "RabbitMQ",
		DestinationType:      newrelic.MessageExchange,
		DestinationName:      string(s.exchange),
		DestinationTemporary: false,
	}
	defer segment.End()
	return s.Publish(bytes, client, key)
}

func NewClientPublisher(conn *rabbitmq.Conn) *StandardPublisher {
	publisher := &StandardPublisher{
		exchange: Client,
		kind:     Topic,
	}
	actual, err := newPublisher(conn, publisher)
	if err != nil {
		panic(err)
	}
	publisher.actual = actual
	return publisher
}

func NewSocialPublisher(conn *rabbitmq.Conn) *StandardPublisher {
	publisher := &StandardPublisher{
		exchange: Social,
		kind:     Topic,
	}
	actual, err := newPublisher(conn, publisher)
	if err != nil {
		panic(err)
	}
	publisher.actual = actual
	return publisher
}

func NewRankingPublisher(conn *rabbitmq.Conn) *StandardPublisher {
	publisher := &StandardPublisher{
		exchange: Ranking,
		kind:     Topic,
	}
	actual, err := newPublisher(conn, publisher)
	if err != nil {
		panic(err)
	}
	publisher.actual = actual
	return publisher
}

func NewGamePublisher(conn *rabbitmq.Conn) *StandardPublisher {
	publisher := &StandardPublisher{
		exchange: Game,
		kind:     Topic,
	}
	actual, err := newPublisher(conn, publisher)
	if err != nil {
		panic(err)
	}
	publisher.actual = actual
	return publisher
}

func newPublisher(conn *rabbitmq.Conn, standard *StandardPublisher, options ...func(*rabbitmq.PublisherOptions)) (*rabbitmq.Publisher, error) {
	options = append(
		options,
		rabbitmq.WithPublisherOptionsExchangeName(string(standard.exchange)),
		rabbitmq.WithPublisherOptionsExchangeKind(string(standard.kind)),
		rabbitmq.WithPublisherOptionsExchangeDurable,
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	return rabbitmq.NewPublisher(conn, options...)
}
