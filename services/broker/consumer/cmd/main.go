package main

import (
	"encoding/json"
	"log"
	"rabbitmq/config"
	"rabbitmq/models"
	"rabbitmq/services/broker/consumer"
	"time"
)

func main() {
	channel, err := consumer.NewRabbitMQ(config.GetConfigForRabbit())

	if err != nil {
		log.Fatalf("error initialize RabbitMQ: %v", err)
	}

	var brokerChannel chan int = make(chan int)
	go func() {
		for msg := range channel {
			var channelInfo models.InputChannel
			if err = json.Unmarshal(msg.Body, &channelInfo); err != nil {
				log.Println(err.Error())
			}
			log.Printf("message retrieved from queue(%s) with youtube_id(%s)\n", config.GetConfigForRabbit().ChannelQueue, channelInfo.YoutubeId)
			time.Sleep(5 * time.Second)
			_ = msg.Nack(false, false)
			log.Printf("message rejected from queue(%s) with youtube_id(%s)\n", config.GetConfigForRabbit().ChannelQueue, channelInfo.YoutubeId)
		}
		brokerChannel <- 1
	}()

	<-brokerChannel
}
