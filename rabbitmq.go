package rabbitmq

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"rabbitmq/config"
)

func NewRabbitMQ(cfg config.RabbitCredentials) (*amqp091.Channel, error) {
	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s", cfg.Username, cfg.Password, cfg.Url)
	connection, err := amqp091.Dial(amqpUrl)
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}
	return channel, err
}
