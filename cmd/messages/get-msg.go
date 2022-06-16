package messages

import (
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/appDAL"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/chatDAL"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/messageDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func GetMsg(ctx *gin.Context) {
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
	chatNumber := ctx.GetInt("chat_number")
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
	msgNumber := ctx.GetInt("msg_number")
	msg, err := messageDAL.GetMessageByNumber(chat.Id, msgNumber)
	if err != nil || msg.Id <= 0 {
		if err != nil {
			log.Err(err).Msg(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the msg"})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no msg with this number"})
		return
	}
	ctx.JSON(http.StatusOK, msg)
}
