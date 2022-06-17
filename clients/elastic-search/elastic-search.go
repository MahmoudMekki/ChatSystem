package esClient

import (
	"fmt"
	"github.com/MahmoudMekki/ChatSystem/config"
	"github.com/olivere/elastic/v7"
	"github.com/rs/zerolog/log"
)

var esClient *elastic.Client

func connectToES() {
	var err error
	esClient, err = elastic.NewClient(elastic.SetURL(config.GetEnvVar("ES_URL")))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	elastic.SetHealthcheck(true)
	elastic.SetSniff(false)
	log.Info().Msg(fmt.Sprintf("Elasticsearch is running on -> %s", config.GetEnvVar("ES_URL")))
}

func GetEsClient() *elastic.Client {
	if esClient == nil {
		connectToES()
	}
	return esClient
}
