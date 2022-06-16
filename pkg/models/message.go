package models

const (
	MessagesTableName = "messages"
)

type Message struct {
	Id      int    `gorm:"column:id;primary_key;auto_increment" json:"-"`
	ChatId  int    `gorm:"column:chat_id;not null" json:"-"`
	Number  int    `gorm:"column:number;not null" json:"number"`
	Content string `gorm:"column:content;not null" json:"content"`
	Chat    Chat   `gorm:"foreign_key:chat_id" json:"-"`
}

func (m *Message) TableName() string {
	return MessagesTableName
}
