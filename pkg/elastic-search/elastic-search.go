package elasticsearch

import (
	"context"
	"encoding/json"
	esClient "github.com/MahmoudMekki/ChatSystem/clients/elastic-search"
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/olivere/elastic/v7"
)

func EsIndex(msgInd models.MessageIndex) error {
	esClient := esClient.GetEsClient()
	jsonData, err := json.Marshal(msgInd)
	if err != nil {
		return err
	}
	_, err = esClient.Index().Index("msgs").BodyJson(string(jsonData)).Do(context.Background())
	return err
}

func AutoComplete(paginator models.Paginator, msgInd models.MessageIndex) (*elastic.SearchResult, error) {
	esClient := esClient.GetEsClient()
	likeQuery := elastic.NewMultiMatchQuery(paginator.GetKeyWord(), "content").Type("phrase_prefix")
	matchChatQuery := elastic.NewMatchQuery("chat_number", msgInd.ChatNumber)
	matchTokenQuery := elastic.NewMatchQuery("app_token", msgInd.AppToken)
	query := elastic.NewBoolQuery().Must(likeQuery, matchChatQuery, matchTokenQuery)
	rslt, err := esClient.Search().
		Index("msgs").
		Query(query).
		From(paginator.GetOffset()).
		Size(paginator.GetLimit()).
		Do(context.Background())
	return rslt, err
}
