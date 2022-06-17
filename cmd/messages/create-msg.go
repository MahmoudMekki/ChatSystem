package messages

import (
	elasticsearch "github.com/MahmoudMekki/ChatSystem/pkg/elastic-search"
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/appDAL"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/chatDAL"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/messageDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func CreateMsg(ctx *gin.Context) {
	token := ctx.GetString("token")
	app, err := appDAL.GetAppByToken(token)
	if err != nil || app.Id <= 0 {
		if app.Id > 0 {
			log.Err(err).Msg(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the chat"})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no app for this token"})
		return
	}
	chatNumber := ctx.GetInt("number")
	chat, err := chatDAL.GetChatByNumber(app.Id, chatNumber)
	if err != nil || chat.Id <= 0 {
		if app.Id > 0 {
			log.Err(err).Msg(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the msg"})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no chats for this number"})
		return
	}
	latestMsgNumber, err := messageDAL.GetMaxNumberOfAppChat(chat.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the msg"})
		return
	}
	msg := models.Message{
		ChatId:  chat.Id,
		Number:  latestMsgNumber + 1,
		Content: ctx.GetString("content"),
	}
	msg, err = messageDAL.CreateMessage(msg)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the msg"})
		return
	}
	msgInd := models.MessageIndex{
		AppToken:      token,
		ChatNumber:    chatNumber,
		MessageNumber: msg.Number,
		Content:       msg.Content,
	}
	if err = elasticsearch.EsIndex(msgInd); err != nil {
		log.Err(err).Msg(err.Error())
	}
	ctx.JSON(http.StatusOK, msg)
}
