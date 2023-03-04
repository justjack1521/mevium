package mevcon

import (
	"fmt"
	"github.com/wagslane/go-rabbitmq"
)

type RabbitMQConfig struct {
	Host     string `required:"true"`
	Username string `required:"true"`
	Password string `required:"true"`
	Port     int    `required:"true"`
}

func (c RabbitMQConfig) getSource() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d", c.Username, c.Password, c.Host, c.Port)
}

func NewRabbitMQConnection(config RabbitMQConfig) (*rabbitmq.Conn, error) {
	conn, err := rabbitmq.NewConn(config.getSource(), rabbitmq.WithConnectionOptionsLogging)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
