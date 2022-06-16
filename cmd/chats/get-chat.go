package chats

import (
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/appDAL"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/chatDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func GetChat(ctx *gin.Context) {
	token := ctx.GetString("token")
	number := ctx.GetInt("number")
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
	chat, err := chatDAL.GetChatByNumber(app.Id, number)
	if err != nil || chat.Id <= 0 {
		if chat.Id > 0 {
			log.Err(err).Msg(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the chat"})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no chat with this number"})
		return
	}
	ctx.JSON(http.StatusOK, chat)
}
