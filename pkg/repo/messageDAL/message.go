package messageDAL

import (
	"github.com/MahmoudMekki/ChatSystem/database"
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
)

func CreateMessage(msg models.Message) (models.Message, error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.Message{}, err
	}
	err = dbConn.Table(models.MessagesTableName).Create(&msg).Error
	if err != nil {
		return models.Message{}, err
	}
	return msg, nil
}
func GetMessageByNumber(chatID, number int) (msg models.Message, err error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.Message{}, err
	}
	err = dbConn.Table(models.MessagesTableName).Where("chat_id=? and number=?", chatID, number).Find(&msg).Error
	if err != nil {
		return models.Message{}, err
	}
	return msg, nil
}
