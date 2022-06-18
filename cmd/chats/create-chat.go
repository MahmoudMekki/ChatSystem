package chats

import (
	"encoding/json"
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/MahmoudMekki/ChatSystem/pkg/rabbit"
	"github.com/MahmoudMekki/ChatSystem/pkg/redis"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func CreateChat(ctx *gin.Context) {
	token := ctx.GetString("token")
	latestChatNumber, err := redis.CacheLastChatNumber(token)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the chat"})
		return
	}
	msg, err := json.Marshal(models.ChatMQMsg{
		ApplicationToken: token,
		Number:           latestChatNumber,
	})
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the chat"})
		return
	}
	err = rabbit.Produce(models.ChatMQTopic, msg)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the chat"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"chat_number": latestChatNumber})
}
