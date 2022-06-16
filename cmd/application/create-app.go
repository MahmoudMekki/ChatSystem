package application

import (
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/appDAL"
	"github.com/MahmoudMekki/ChatSystem/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func CreateApp(ctx *gin.Context) {
	name, _ := ctx.GetQuery("name")
	token, err := utils.GenerateRandomString()
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while creating the application"})
		return
	}
	app := models.Application{Name: name, Token: token}
	app, err = appDAL.CreateApp(app)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while creating the application"})
		return
	}
	ctx.JSON(http.StatusOK, app)
}
