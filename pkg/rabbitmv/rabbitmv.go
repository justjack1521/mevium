package rabbitmv

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

var (
	errClientIDNotFound                 = errors.New("client id not found")
	errExtractClientIDFromMessageHeader = func(err error) error {
		return fmt.Errorf("failed to extract client id from message header: %w", err)
	}
)

type ExchangeKind string

const (
	Direct  ExchangeKind = "direct"
	Fanout  ExchangeKind = "fanout"
	Headers ExchangeKind = "headers"
	Topic   ExchangeKind = "topic"
)

type Exchange string

const (
	Client  Exchange = "client"
	Social  Exchange = "social"
	Ranking Exchange = "ranking"
)

type RoutingKey string

const (
	ClientNotification RoutingKey = "client.notification"
	ClientHeartbeat    RoutingKey = "client.heartbeat"
	ClientConnected    RoutingKey = "client.connected"
	ClientDisconnected RoutingKey = "client.disconnected"
	PlayerComment      RoutingKey = "player.comment"
	PlayerCompanion    RoutingKey = "player.companion"
	PlayerPosition     RoutingKey = "player.position"
	PlayerPresence     RoutingKey = "player.presence"
	BattleComplete     RoutingKey = "battle.complete"
)

type Queue string

const (
	ClientUpdate    Queue = "client.update"
	CommentUpdate   Queue = "comment.update"
	CompanionUpdate Queue = "companion.update"
	PositionUpdate  Queue = "position.update"
	PresenceUpdate  Queue = "presence.update"
	SocialUpdate    Queue = "social.update"
	RankingUpdate   Queue = "ranking.update"
)

func NewConsumer(conn *rabbitmq.Conn, handler rabbitmq.Handler, queue Queue, key RoutingKey, exchange Exchange, options ...func(*rabbitmq.ConsumerOptions)) (*rabbitmq.Consumer, error) {
	options = append(options, rabbitmq.WithConsumerOptionsRoutingKey(string(key)), rabbitmq.WithConsumerOptionsExchangeName(string(exchange)))
	return rabbitmq.NewConsumer(conn, handler, string(queue), options...)
}

func NewPublisher(conn *rabbitmq.Conn, exchange Exchange, kind ExchangeKind, options ...func(*rabbitmq.PublisherOptions)) (*rabbitmq.Publisher, error) {
	options = append(
		options,
		rabbitmq.WithPublisherOptionsExchangeName(string(exchange)),
		rabbitmq.WithPublisherOptionsExchangeKind(string(kind)),
		rabbitmq.WithPublisherOptionsExchangeDurable,
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	return rabbitmq.NewPublisher(conn, options...)
}

func ExtractClientID(d rabbitmq.Delivery) (uuid.UUID, error) {
	id, ok := d.Headers["client_id"]
	if ok == false {
		return uuid.Nil, errExtractClientIDFromMessageHeader(errClientIDNotFound)
	}
	client, err := uuid.FromString(id.(string))
	if err != nil {
		return client, errExtractClientIDFromMessageHeader(err)
	}
	return client, nil
}
func ExtractPlayerID(d rabbitmq.Delivery) (uuid.UUID, error) {
	id, ok := d.Headers["player_id"]
	if ok == false {
		return uuid.Nil, errExtractClientIDFromMessageHeader(errClientIDNotFound)
	}
	client, err := uuid.FromString(id.(string))
	if err != nil {
		return client, errExtractClientIDFromMessageHeader(err)
	}
	return client, nil
}

func ClientPublishingTable(client uuid.UUID) rabbitmq.Table {
	return rabbitmq.Table{
		"client_id": client.String(),
	}
}

func PlayerPublishingTable(client uuid.UUID) rabbitmq.Table {
	return rabbitmq.Table{
		"player_id": client.String(),
	}
}

type ConsumerHandler func(d rabbitmq.Delivery) (rabbitmq.Action, error)

func RabbitConsumerMiddleware(handler ConsumerHandler) rabbitmq.Handler {
	return func(d rabbitmq.Delivery) rabbitmq.Action {
		result, _ := handler(d)
		return result
	}
}

func RabbitConsumeLoggerMiddleWare(logger *logrus.Logger, handler ConsumerHandler) ConsumerHandler {
	return func(d rabbitmq.Delivery) (rabbitmq.Action, error) {
		logger.WithFields(logrus.Fields{
			"exchange":    d.Exchange,
			"routing_key": d.RoutingKey,
		}).Info("Message Received for Consumption")

		action, err := handler(d)

		defer func() {
			if err != nil {
				logger.WithFields(logrus.Fields{
					"exchange":    d.Exchange,
					"routing_key": d.RoutingKey,
					"action":      action,
				}).WithError(err).Error("Failed to Consume Message")
			} else {
				logger.WithFields(logrus.Fields{
					"exchange":    d.Exchange,
					"routing_key": d.RoutingKey,
					"action":      action,
				}).Info("Message Consumed")
			}
		}()
		return action, err
	}
}
