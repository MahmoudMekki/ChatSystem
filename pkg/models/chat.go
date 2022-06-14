package models

const (
	ChatTableName = "chats"
)

type Chat struct {
	Id            int         `gorm:"column:id;primary_key;auto_increment" json:"_"`
	ApplicationId int         `gorm:"column:app_id;not null" json:"application_id"`
	Number        int         `gorm:"column:number;not null" json:"number"`
	MessageCount  int         `gorm:"column:msg_count;default:0" json:"message_count"`
	Application   Application `gorm:"foreignkey:app_id" json:"application"`
}

func (c *Chat) TableName() string {
	return ChatTableName
}
