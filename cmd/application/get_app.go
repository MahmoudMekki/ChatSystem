package application

import (
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/appDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func GetApp(ctx *gin.Context) {
	token := ctx.GetString("token")
	app, err := appDAL.GetAppByToken(token)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while retrieving the app"})
		return
	}
	ctx.JSON(http.StatusOK, app)
}
