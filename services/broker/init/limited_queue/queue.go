package limited_queue

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

	/*
		Configuration:
		Max message count: 5,
		Queue time to live: 30 seconds
		Message time to live: 10 seconds
	*/

	args := amqp091.Table{
		amqp091.QueueTypeArg:       amqp091.QueueTypeClassic,
		amqp091.QueueMaxLenArg:     5,
		amqp091.QueueTTLArg:        30000,
		amqp091.QueueMessageTTLArg: 10000,
	}

	_, err = channel.QueueDeclare(
		cfg.ChannelQueue,
		true,
		false,
		false,
		false,
		args,
	)

	if err != nil {
		return nil, err
	}

	err = channel.Qos(1, 0, false)
	if err != nil {
		return nil, err
	}

	return channel, nil
}
