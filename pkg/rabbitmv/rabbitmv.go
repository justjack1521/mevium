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

func RabbitConsumeLoggerMiddleWare(logger *logrus.Logger, handler rabbitmq.Handler) rabbitmq.Handler {
	return func(d rabbitmq.Delivery) rabbitmq.Action {

		logger.WithFields(logrus.Fields{
			"exchange":    d.Exchange,
			"routing_key": d.RoutingKey,
		}).Info("Message Received")

		action := handler(d)

		defer func() {
			logger.WithFields(logrus.Fields{
				"exchange":    d.Exchange,
				"routing_key": d.RoutingKey,
				"action":      action,
			}).Info("Message Response")
		}()

		return action
	}
}
