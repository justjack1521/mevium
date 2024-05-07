package rabbitmv

import (
	"context"
	"github.com/justjack1521/mevrpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type ConsumerContext struct {
	context.Context
	Transaction *newrelic.Transaction
	UserID      uuid.UUID
	PlayerID    uuid.UUID
	Delivery    rabbitmq.Delivery
}

type ConsumerHandler func(ctx *ConsumerContext) (rabbitmq.Action, error)

type StandardConsumer struct {
	actual     *rabbitmq.Consumer
	relic      *newrelic.Application
	Queue      Queue
	RoutingKey RoutingKey
	Exchange   Exchange
	closed     bool
}

func (s *StandardConsumer) Close() {
	if s.closed == true {
		return
	}
	s.closed = true
	s.actual.Close()
}

func (s *StandardConsumer) WithNewRelic(r *newrelic.Application) *StandardConsumer {
	s.relic = r
	return s
}

func NewStandardConsumer(conn *rabbitmq.Conn, queue Queue, key RoutingKey, exchange Exchange, handler ConsumerHandler, opts ...func(*rabbitmq.ConsumerOptions)) *StandardConsumer {

	consumer := &StandardConsumer{
		Queue:      queue,
		RoutingKey: key,
		Exchange:   exchange,
	}

	var options = []func(*rabbitmq.ConsumerOptions){
		rabbitmq.WithConsumerOptionsRoutingKey(string(key)),
		rabbitmq.WithConsumerOptionsExchangeName(string(exchange)),
	}

	opts = append(opts, options...)

	actual, err := rabbitmq.NewConsumer(
		conn,
		consumer.StandardConsumption(handler),
		string(queue),
		opts...,
	)

	if err != nil {
		panic(err)
	}

	consumer.actual = actual
	return consumer
}

func (s *StandardConsumer) StandardConsumption(handler ConsumerHandler) rabbitmq.Handler {
	return func(d rabbitmq.Delivery) rabbitmq.Action {

		user, err := ExtractUserID(d)
		if err != nil {
			return rabbitmq.NackDiscard
		}

		player, err := ExtractPlayerID(d)
		if err != nil {
			return rabbitmq.NackDiscard
		}

		ctx := &ConsumerContext{
			Context:  mevrpc.NewOutgoingContext(context.Background(), user, player),
			Delivery: d,
			UserID:   user,
			PlayerID: player,
		}

		if s.relic != nil {
			ctx.Transaction = s.relic.StartTransaction("message." + d.RoutingKey + ":" + d.Exchange)
			//ctx.Context = newrelic.NewContext(ctx.Context, ctx.Transaction)
			ctx.Transaction.AddAttribute("user.id", ctx.UserID.String())
			ctx.Transaction.AddAttribute("player.id", ctx.PlayerID.String())
			ctx.Transaction.AddAttribute("message.routingKey", d.RoutingKey)
			ctx.Transaction.AddAttribute("message.exchange", d.Exchange)
			ctx.Transaction.AddAttribute("message.type", d.Type)
			defer ctx.Transaction.End()
		}
		result, _ := handler(ctx)
		return result
	}
}

func ConsumeLoggerMiddleWare(logger *logrus.Logger, handler ConsumerHandler) ConsumerHandler {
	return func(ctx *ConsumerContext) (rabbitmq.Action, error) {
		logger.WithFields(logrus.Fields{
			"exchange":    ctx.Delivery.Exchange,
			"routing_key": ctx.Delivery.RoutingKey,
		}).Info("Message Received for Consumption")

		action, err := handler(ctx)

		defer func() {
			if err != nil {
				logger.WithFields(logrus.Fields{
					"exchange":    ctx.Delivery.Exchange,
					"routing_key": ctx.Delivery.RoutingKey,
					"action":      action,
				}).WithError(err).Error("Failed to Consume Message")
			} else {
				logger.WithFields(logrus.Fields{
					"exchange":    ctx.Delivery.Exchange,
					"routing_key": ctx.Delivery.RoutingKey,
					"action":      action,
				}).Info("Message Consumed")
			}
		}()
		return action, err
	}
}

func ExtractUserID(d rabbitmq.Delivery) (uuid.UUID, error) {
	id, ok := d.Headers[userIDHeaderKey]
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
	id, ok := d.Headers[playerIDHeaderKey]
	if ok == false {
		return uuid.Nil, errExtractClientIDFromMessageHeader(errClientIDNotFound)
	}
	client, err := uuid.FromString(id.(string))
	if err != nil {
		return client, errExtractClientIDFromMessageHeader(err)
	}
	return client, nil
}
