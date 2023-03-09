package consumer

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"rabbitmq/config"
)

func NewRabbitMQ(cfg config.RabbitCredentials) (<-chan amqp091.Delivery, error) {
	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s", cfg.Username, cfg.Password, cfg.Url)
	connection, err := amqp091.Dial(amqpUrl)
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	//queue, err := channel.QueueDeclare(cfg.ChannelQueue, true, false, false, false, nil)
	//if err != nil {
	//	return nil, err
	//}

	err = channel.Qos(1, 0, false)
	if err != nil {
		return nil, err
	}

	messageChannel, err := channel.Consume(
		cfg.ChannelQueue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return messageChannel, nil

}
