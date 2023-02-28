package rabbitmv

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
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

func ClientPublishingTable(client uuid.UUID) rabbitmq.Table {
	return rabbitmq.Table{
		"client_id": client.String(),
	}
}
