package main

import (
	"encoding/json"
	"github.com/MahmoudMekki/ChatSystem/bin/crons"
	"github.com/MahmoudMekki/ChatSystem/clients/rabbitMQ"
	"github.com/MahmoudMekki/ChatSystem/database"
	elasticsearch "github.com/MahmoudMekki/ChatSystem/pkg/elastic-search"
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/appDAL"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/chatDAL"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/messageDAL"
	"github.com/jasonlvhit/gocron"
	"github.com/rs/zerolog/log"
)

func init() {
	err := database.CreateDBConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	rabbitMQ.GetRabbitMQCConsumeChannel()
}
func main() {
	channel := rabbitMQ.GetRabbitMQCConsumeChannel()
	messages, err := channel.Consume(
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
	}
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
	go func() {
		for message := range messages {
			var msgMq models.MessageMQMsg
			err := json.Unmarshal(message.Body, &msgMq)
			if err != nil {
				log.Err(err).Msg(err.Error())
				continue
			}
			app, err := appDAL.GetAppByToken(msgMq.ApplicationToken)
			if err != nil || app.Id <= 0 {
				if app.Id > 0 {
					log.Err(err).Msg(err.Error())
					continue
				}
				log.Info().Msg("no app for this token")
				continue
			}
			chat, err := chatDAL.GetChatByNumber(app.Id, msgMq.ChatNumber)
			if err != nil || chat.Id <= 0 {
				if app.Id > 0 {
					log.Err(err).Msg(err.Error())
					continue
				}
				log.Info().Msg("no chat for this number")
				continue
			}
			msg := models.Message{
				ChatId:  chat.Id,
				Number:  msgMq.MsgNumber,
				Content: msgMq.Content,
			}
			msg, err = messageDAL.CreateMessage(msg)
			if err != nil {
				log.Err(err).Msg(err.Error())
				continue
			}
			msgInd := models.MessageIndex{
				AppToken:      msgMq.ApplicationToken,
				ChatNumber:    msgMq.ChatNumber,
				MessageNumber: msgMq.MsgNumber,
				Content:       msgMq.Content,
			}
			if err = elasticsearch.EsIndex(msgInd); err != nil {
				log.Err(err).Msg(err.Error())
			}
		}
	}()
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
		}
	}()
	go func() {
		gocron.Every(30).Minutes().Do(crons.UpdateAppChatCount)
		gocron.Every(35).Minutes().Do(crons.UpdateAppChatCount)
		<-gocron.Start()
	}()
	<-forever
}
