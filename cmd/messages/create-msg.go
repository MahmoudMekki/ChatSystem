package messages

import (
	"encoding/json"
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/MahmoudMekki/ChatSystem/pkg/rabbit"
	"github.com/MahmoudMekki/ChatSystem/pkg/redis"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func CreateMsg(ctx *gin.Context) {
	token := ctx.GetString("token")
	chatNumber := ctx.GetInt("chat_number")
	latestMsgNumber, err := redis.CacheLastMsgNumber(token, chatNumber)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the chat"})
		return
	}
	mqMsg, err := json.Marshal(models.MessageMQMsg{
		ApplicationToken: token,
		ChatNumber:       chatNumber,
		MsgNumber:        latestMsgNumber,
		Content:          ctx.GetString("content"),
	})
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the chat"})
		return
	}
	err = rabbit.Produce(models.MessagesMQTopic, mqMsg)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating the chat"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg_number": latestMsgNumber})
}
