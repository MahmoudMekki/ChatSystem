package messages

import (
	"encoding/json"
	elasticsearch "github.com/MahmoudMekki/ChatSystem/pkg/elastic-search"
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func AutoComplete(ctx *gin.Context) {
	paginator := models.Paginator{
		Limit:   ctx.GetInt("limit"),
		Page:    ctx.GetInt("page"),
		KeyWord: ctx.GetString("keyword"),
	}
	msgInd := models.MessageIndex{
		AppToken:   ctx.GetString("token"),
		ChatNumber: ctx.GetInt("chat_number"),
	}

	rslt, err := elasticsearch.AutoComplete(paginator, msgInd)
	if err != nil {
		log.Err(err).Msg(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while requesting the data"})
		return
	}
	var messages []models.MessageIndex
	for _, v := range rslt.Hits.Hits {
		var msg models.MessageIndex
		err := json.Unmarshal(v.Source, &msg)
		if err != nil {
			log.Err(err).Msg(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while requesting the data"})
			return
		}
		messages = append(messages, msg)
	}
	ctx.Header("X-TOTAL-HITS", strconv.Itoa(int(rslt.Hits.TotalHits.Value)))
	ctx.JSON(http.StatusOK, messages)
}
