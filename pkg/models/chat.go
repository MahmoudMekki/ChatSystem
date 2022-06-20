package models

const (
	ChatTableName = "chats"
	ChatMQTopic   = "chats"
)

type Chat struct {
	Id            int       `gorm:"column:id;primary_key;auto_increment" json:"-"`
	ApplicationId int       `gorm:"column:app_id;not null" json:"-"`
	Number        int       `gorm:"column:number;not null" json:"number"`
	MessageCount  int       `gorm:"column:msg_count;default:0" json:"message_count"`
	Messages      []Message `json:"messages,omitempty"`
}
type ChatMQMsg struct {
	ApplicationToken string `json:"token"`
	Number           int    `json:"number"`
}

func (c *Chat) TableName() string {
	return ChatTableName
}
