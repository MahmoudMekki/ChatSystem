package rabbitMQ

import (
	"github.com/MahmoudMekki/ChatSystem/config"
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

var publishChannel *amqp.Channel
var consumeChannel *amqp.Channel

func establishRabbitMQ() {
	var err error
	rabbitConnection, err := amqp.Dial(config.GetEnvVar("RABBITMQ_URL"))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	publishChannel, err = rabbitConnection.Channel()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	_, err = publishChannel.QueueDeclare(models.ChatMQTopic, true, false, false, false, nil)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	_, err = publishChannel.QueueDeclare(models.ChatMQTopic, true, false, false, false, nil)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	consumeChannel, err = rabbitConnection.Channel()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func GetRabbitMQCPublishChannel() *amqp.Channel {
	if publishChannel == nil {
		establishRabbitMQ()
	}
	return publishChannel
}
func GetRabbitMQCConsumeChannel() *amqp.Channel {
	if consumeChannel == nil {
		establishRabbitMQ()
	}
	return consumeChannel
}
