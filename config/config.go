package config

import "rabbitmq/constants"

type RabbitCredentials struct {
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Url          string `mapstructure:"url"`
	ChannelQueue string `mapstructure:"channelQueue"`
}

func GetConfigForRabbit() RabbitCredentials {
	return RabbitCredentials{
		Username:     "guest",
		Password:     "guest",
		Url:          "localhost:5672",
		ChannelQueue: constants.RabbitQueueName,
	}
}
