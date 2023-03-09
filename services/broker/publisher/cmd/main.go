package main

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq"
	"rabbitmq/config"
	"rabbitmq/constants"
	"rabbitmq/models"
	"strconv"
	"time"
)

func main() {
	amqpChannel, err := rabbitmq.NewRabbitMQ(config.GetConfigForRabbit())

	defer amqpChannel.Close()

	if err != nil {
		log.Fatalf("error when connecting to rabbit: %v\n", err)
	}

	//var uuidObj uuid.UUID
	//if uuidObj, err = uuid.NewUUID(); err != nil {
	//	log.Fatalf("error with getting uuid: %v\n", err)
	//}

	var i int64
	for i = 1; i <= 30; i++ {
		err = Publish(amqpChannel, models.InputChannel{
			YoutubeId: strconv.FormatInt(i, 10),
			IsForeign: false,
		}, constants.RabbitQueueName)

		if err != nil {
			log.Printf("error when publish to channel: %v\n", err)
		}
	}

}

func Publish(amqpChannel *amqp091.Channel, input models.InputChannel, queueName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body, err := json.Marshal(input)
	if err != nil {
		log.Println(err.Error())
	}
	err = amqpChannel.PublishWithContext(ctx, "", queueName, false, false,
		amqp091.Publishing{
			DeliveryMode: amqp091.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		log.Println("error to publish message")
		return err
	}
	log.Printf("youtube channel with youtube_id = %s published to %s successfully", input.YoutubeId, queueName)
	return nil
}
