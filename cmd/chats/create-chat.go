package chats

import (
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/appDAL"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/chatDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func CreateChat(ctx *gin.Context) {
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
	latestChatNumber, err := chatDAL.GetMaxNumberOfAppChat(app.Id)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the chat"})
		return
	}
	chat := models.Chat{
		ApplicationId: app.Id,
		Number:        latestChatNumber + 1,
	}
	chat, err = chatDAL.CreateChat(chat)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the chat"})
		return
	}
	ctx.JSON(http.StatusOK, chat)
}
