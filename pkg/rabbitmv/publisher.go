package rabbitmv

import (
	"context"
	"github.com/newrelic/go-agent/v3/newrelic"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type StandardPublisher struct {
	exchange Exchange
	kind     ExchangeKind
	closed   bool
	actual   *rabbitmq.Publisher
	logger   *logrus.Logger
}

func (s *StandardPublisher) AttachLogger(logger *logrus.Logger) {
	s.logger = logger
}

func (s *StandardPublisher) Close() {
	if s.closed == true {
		return
	}
	s.closed = true
	s.actual.Close()
}

func (s *StandardPublisher) Publish(bytes []byte, user, player uuid.UUID, key RoutingKey, options ...func(*rabbitmq.PublishOptions)) error {
	options = append(
		options,
		rabbitmq.WithPublishOptionsExchange(string(s.exchange)),
		rabbitmq.WithPublishOptionsHeaders(IdentityPublishingTable(user, player)),
	)
	if err := s.actual.Publish(bytes, []string{string(key)}, options...); err != nil {
		if s.logger != nil {
			s.logger.WithFields(logrus.Fields{
				"exchange":    s.exchange,
				"routing.key": key,
				"length":      len(bytes),
			}).WithError(err).Error("Message Failed")
		}
		return err
	}
	if s.logger != nil {
		s.logger.WithFields(logrus.Fields{
			"exchange":    s.exchange,
			"routing.key": key,
			"length":      len(bytes),
		}).Info("Message Published")
	}
	return nil
}

func (s *StandardPublisher) PublishWithContext(ctx context.Context, bytes []byte, user, player uuid.UUID, key RoutingKey) error {
	txn := newrelic.FromContext(ctx)
	return s.PublishWithSegment(txn, bytes, user, player, key)
}

func (s *StandardPublisher) PublishWithSegment(txn *newrelic.Transaction, bytes []byte, user, player uuid.UUID, key RoutingKey) error {
	segment := newrelic.MessageProducerSegment{
		StartTime:            txn.StartSegmentNow(),
		Library:              "RabbitMQ",
		DestinationType:      newrelic.MessageExchange,
		DestinationName:      string(s.exchange),
		DestinationTemporary: false,
	}
	defer segment.End()
	return s.Publish(bytes, user, player, key)
}

func NewClientPublisher(conn *rabbitmq.Conn, options ...func(publisherOptions *rabbitmq.PublisherOptions)) *StandardPublisher {
	publisher := &StandardPublisher{
		exchange: Client,
		kind:     Topic,
	}
	actual, err := newPublisher(conn, publisher, options...)
	if err != nil {
		panic(err)
	}
	publisher.actual = actual
	return publisher
}

func NewSocialPublisher(conn *rabbitmq.Conn, options ...func(publisherOptions *rabbitmq.PublisherOptions)) *StandardPublisher {
	publisher := &StandardPublisher{
		exchange: Social,
		kind:     Topic,
	}
	actual, err := newPublisher(conn, publisher, options...)
	if err != nil {
		panic(err)
	}
	publisher.actual = actual
	return publisher
}

func NewRankingPublisher(conn *rabbitmq.Conn, options ...func(publisherOptions *rabbitmq.PublisherOptions)) *StandardPublisher {
	publisher := &StandardPublisher{
		exchange: Ranking,
		kind:     Topic,
	}
	actual, err := newPublisher(conn, publisher, options...)
	if err != nil {
		panic(err)
	}
	publisher.actual = actual
	return publisher
}

func NewGamePublisher(conn *rabbitmq.Conn, options ...func(publisherOptions *rabbitmq.PublisherOptions)) *StandardPublisher {
	publisher := &StandardPublisher{
		exchange: Game,
		kind:     Topic,
	}
	actual, err := newPublisher(conn, publisher, options...)
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

func IdentityPublishingTable(user, player uuid.UUID) rabbitmq.Table {
	return rabbitmq.Table{
		userIDHeaderKey:   user.String(),
		playerIDHeaderKey: player.String(),
	}
}
