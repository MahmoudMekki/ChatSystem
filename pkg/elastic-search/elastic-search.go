package elasticsearch

import (
	"context"
	"encoding/json"
	esClient "github.com/MahmoudMekki/ChatSystem/clients/elastic-search"
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/olivere/elastic/v7"
)

func EsIndex(msg models.Message) error {
	esClient := esClient.GetEsClient()
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = esClient.Index().Index("msgs").BodyJson(string(jsonData)).Do(context.Background())
	return err
}

func AutoComplete(keyword string, chatID int) ([]models.Message, error) {
	esClient := esClient.GetEsClient()
	likeQuery := elastic.NewMultiMatchQuery(keyword, "content").Type("phrase_prefix")
	matchQuery := elastic.NewMatchQuery("ChatId", chatID)
	query := elastic.NewBoolQuery().Must(likeQuery, matchQuery)
	rslt, err := esClient.Search().Index("msgs").Query(query).Do(context.Background())
	if err != nil {
		return nil, err
	}
	var messages []models.Message
	for _, v := range rslt.Hits.Hits {
		var msg models.Message
		err := json.Unmarshal(v.Source, &msg)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
