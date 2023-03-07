package rabbitmv

import (
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type ConsumerHandler func(d rabbitmq.Delivery) (rabbitmq.Action, error)

type StandardConsumer struct {
	actual     *rabbitmq.Consumer
	Queue      Queue
	RoutingKey RoutingKey
	Exchange   Exchange
}

func (s *StandardConsumer) Close() {
	s.actual.Close()
}

func NewStandardConsumer(conn *rabbitmq.Conn, queue Queue, key RoutingKey, exchange Exchange, handler ConsumerHandler) *StandardConsumer {

	consumer := &StandardConsumer{
		Queue:      queue,
		RoutingKey: key,
		Exchange:   exchange,
	}

	var options = []func(*rabbitmq.ConsumerOptions){
		rabbitmq.WithConsumerOptionsRoutingKey(string(key)),
		rabbitmq.WithConsumerOptionsExchangeName(string(exchange)),
	}

	actual, err := rabbitmq.NewConsumer(
		conn,
		consumer.StandardConsumption(handler),
		string(queue),
		options...,
	)

	if err != nil {
		panic(err)
	}

	consumer.actual = actual
	return consumer
}

func (s *StandardConsumer) StandardConsumption(handler ConsumerHandler) rabbitmq.Handler {
	return func(d rabbitmq.Delivery) rabbitmq.Action {
		result, _ := handler(d)
		return result
	}
}

func ConsumeLoggerMiddleWare(logger *logrus.Logger, handler ConsumerHandler) ConsumerHandler {
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
