package rabbit

import (
	"github.com/MahmoudMekki/ChatSystem/clients/rabbitMQ"
	"github.com/streadway/amqp"
)

func Produce(topic string, msg []byte) error {
	channel := rabbitMQ.GetRabbitMQCPublishChannel()
	message := amqp.Publishing{
		ContentType: "json/application",
		Body:        msg,
	}
	err := channel.Publish("", topic, true, false, message)
	return err
}
