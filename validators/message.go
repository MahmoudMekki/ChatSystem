package validators

import (
	"github.com/MahmoudMekki/ChatSystem/config"
	"github.com/faceair/jio"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ValidateCreateMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, ok := ctx.Params.Get("token")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		size := config.GetEnvVar("TOKEN_SIZE")
		tokLength, err := strconv.Atoi(size)
		if err != nil || len(token) != tokLength {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
			return
		}
		chatNumStr, ok := ctx.Params.Get("chat_number")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request,chat number is required"})
			return
		}
		chatNum, err := strconv.Atoi(chatNumStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request,invalid chat number"})
			return
		}
		jsonData, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		data, err := jio.ValidateJSON(&jsonData, jio.Object().Keys(jio.K{
			"content": jio.String().Required(),
		}))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.Set("token", token)
		ctx.Set("chat_number", chatNum)
		ctx.Set("content", data["content"])
		ctx.Next()
	}
}
func ValidateGetMessage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, ok := ctx.Params.Get("token")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		size := config.GetEnvVar("TOKEN_SIZE")
		tokLength, err := strconv.Atoi(size)
		if err != nil || len(token) != tokLength {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
			return
		}
		chatNumStr, ok := ctx.Params.Get("chat_number")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request,chat number is required"})
			return
		}
		chatNum, err := strconv.Atoi(chatNumStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request,invalid chat number"})
			return
		}
		msgNumStr, ok := ctx.Params.Get("msg_number")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request, msg number is required"})
			return
		}
		msgNum, err := strconv.Atoi(msgNumStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request,invalid msg number"})
			return
		}
		ctx.Set("token", token)
		ctx.Set("chat_number", chatNum)
		ctx.Set("msg_number", msgNum)
		ctx.Next()
	}
}

func ValidateMsgAutoComplete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, ok := ctx.Params.Get("token")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		size := config.GetEnvVar("TOKEN_SIZE")
		tokLength, err := strconv.Atoi(size)
		if err != nil || len(token) != tokLength {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
			return
		}
		chatNumStr, ok := ctx.Params.Get("chat_number")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request,chat number is required"})
			return
		}
		chatNum, err := strconv.Atoi(chatNumStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request,invalid chat number"})
			return
		}
		keyword, _ := ctx.GetQuery("keyword")
		limitStr, _ := ctx.GetQuery("limit")
		pageStr, _ := ctx.GetQuery("page")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = 10
		}
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
		ctx.Set("token", token)
		ctx.Set("chat_number", chatNum)
		ctx.Set("limit", limit)
		ctx.Set("page", page)
		ctx.Set("keyword", keyword)
		ctx.Next()
	}
}
