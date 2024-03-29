package models

const (
	MessagesTableName = "messages"
	MessagesMQTopic   = "messages"
)

type Message struct {
	Id      int    `gorm:"column:id;primary_key;auto_increment" json:"-"`
	ChatId  int    `gorm:"column:chat_id;not null" json:"-"`
	Number  int    `gorm:"column:number;not null" json:"number"`
	Content string `gorm:"column:content;not null" json:"content"`
	Chat    Chat   `gorm:"foreign_key:chat_id" json:"-"`
}
type MessageIndex struct {
	AppToken      string `json:"app_token"`
	ChatNumber    int    `json:"chat_number"`
	MessageNumber int    `json:"message_number"`
	Content       string `json:"content"`
}
type MessageMQMsg struct {
	ApplicationToken string `json:"token"`
	ChatNumber       int    `json:"chat_number"`
	MsgNumber        int    `json:"msg_number"`
	Content          string `json:"content"`
}

func (m *Message) TableName() string {
	return MessagesTableName
}
