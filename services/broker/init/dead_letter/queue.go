package dead_letter

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
		Message time to live: 10 seconds
	*/

	args := amqp091.Table{
		amqp091.QueueTypeArg:   amqp091.QueueTypeClassic,
		amqp091.QueueMaxLenArg: 5,
		//amqp091.QueueTTLArg:         30000,
		amqp091.QueueMessageTTLArg:  10000,
		"x-dead-letter-exchange":    "dlx_exchange",
		"x-dead-letter-routing-key": "dlx_key",
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

	/*
		Specify the exchange normally and declare it as a backup for a queue:
	*/

	err = channel.ExchangeDeclare(
		"dlx_exchange",
		"direct",
		true,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = createAndBindDeadLetterQueue(channel)
	if err != nil {
		return nil, err
	}

	err = channel.Qos(1, 0, false)
	if err != nil {
		return nil, err
	}

	return channel, nil
}

func createAndBindDeadLetterQueue(channel *amqp091.Channel) error {
	_, err := channel.QueueDeclare(
		"dead_letter_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = channel.QueueBind(
		"dead_letter_queue",
		"dlx_key",
		"dlx_exchange",
		false,
		nil,
	)
	return err
}
