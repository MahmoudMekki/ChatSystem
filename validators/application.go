package validators

import (
	"github.com/MahmoudMekki/ChatSystem/config"
	"github.com/faceair/jio"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ValidateCreateUpdateApplication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jsonData, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		data, err := jio.ValidateJSON(&jsonData, jio.Object().Keys(jio.K{
			"name": jio.String().Min(3).Max(10).Required(),
		}))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.Set("name", data["name"])
		ctx.Next()
	}
}

func ValidateGetApplication() gin.HandlerFunc {
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
