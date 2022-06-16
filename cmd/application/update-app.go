package application

import (
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/appDAL"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func UpdateApp(ctx *gin.Context) {
	name := ctx.GetString("name")
	token := ctx.Param("token")
	app := models.Application{Name: name, Token: token}
	app, err := appDAL.UpdateApp(app)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating the app"})
		return
	}
	if app.Id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no apps for the provided token"})
		return
	}
	ctx.JSON(http.StatusOK, app)
}
