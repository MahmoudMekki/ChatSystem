package main

import (
	"encoding/json"
	"fmt"
	"github.com/MahmoudMekki/ChatSystem/clients/rabbitMQ"
	"github.com/MahmoudMekki/ChatSystem/database"
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/appDAL"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/chatDAL"
	"github.com/rs/zerolog/log"
)

func init() {
	err := database.CreateDBConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
func main() {
	channel := rabbitMQ.GetRabbitMQCConsumeChannel()
	/*	messages, err := channel.Consume(
			models.MessagesMQTopic,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}*/
	chats, err := channel.Consume(
		models.ChatMQTopic,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	forever := make(chan bool)
	/*go func() {
		for message := range messages {

		}
	}()*/
	go func() {
		for chat := range chats {
			var chatMsg models.ChatMQMsg
			err := json.Unmarshal(chat.Body, &chatMsg)
			if err != nil {
				log.Err(err).Msg(err.Error())
				continue
			}
			app, err := appDAL.GetAppByToken(chatMsg.ApplicationToken)
			if err != nil || app.Id <= 0 {
				if app.Id > 0 {
					log.Err(err).Msg(err.Error())
					continue
				}
				log.Info().Msg("No app for this token")
				continue
			}
			chatInternal := models.Chat{
				ApplicationId: app.Id,
				Number:        chatMsg.Number,
			}
			chatInternal, err = chatDAL.CreateChat(chatInternal)
			if err != nil {
				log.Err(err).Msg(err.Error())
				continue
			}
			log.Info().Msg(fmt.Sprintf("task with ID %s is successfully done", chat.MessageId))
		}
	}()
	<-forever
}
