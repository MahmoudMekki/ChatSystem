package chatDAL

import (
	"github.com/MahmoudMekki/ChatSystem/database"
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
)

func CreateChat(chat models.Chat) (models.Chat, error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.Chat{}, err
	}
	err = dbConn.Table(models.ChatTableName).Create(&chat).Error
	if err != nil {
		return models.Chat{}, err
	}
	return chat, nil
}
func GetChatByNumber(appID, number int) (chat models.Chat, err error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.Chat{}, err
	}
	err = dbConn.Preload("Messages").Table(models.ChatTableName).Where("app_id=? and number=?", appID, number).Find(&chat).Error
	if err != nil {
		return models.Chat{}, err
	}
	return chat, nil
}

func GetMaxNumberOfAppChat(appID int) (int, error) {
	var chat models.Chat
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return 0, err
	}
	err = dbConn.Raw(`select * from chats where app_id =? and number = (select MAX(number) from chats where app_id=?)`, appID, appID).Scan(&chat).Error
	if err != nil {
		return 0, err
	}
	return chat.Number, nil
}