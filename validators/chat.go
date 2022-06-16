package validators

import (
	"github.com/MahmoudMekki/ChatSystem/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ValidateCreateChat() gin.HandlerFunc {
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
		ctx.Set("token", token)
		ctx.Next()
	}
}
func ValidateGetChat() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, ok := ctx.Params.Get("token")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		numberStr, ok := ctx.Params.Get("chat_number")
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
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Chat number"})
			return
		}
		ctx.Set("token", token)
		ctx.Set("number", number)
		ctx.Next()
	}
}
