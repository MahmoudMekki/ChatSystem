package validators

import (
	"github.com/faceair/jio"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func ValidateCreateApplication() gin.HandlerFunc {
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
