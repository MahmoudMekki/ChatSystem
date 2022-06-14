package models

import "time"

const (
	MessagesTableName = "messages"
)

type Message struct {
	Id      int       `gorm:"column:id;primary_key" json:"-"`
	ChatId  int       `gorm:"column:chat_id;not null" json:"chat_id"`
	Number  int       `gorm:"column:number;not null" json:"number"`
	Content time.Time `gorm:"column:content;not null" json:"content"`
	Chat    Chat      `gorm:"foreignkey:chat_id" json:"chat""`
}

func (m *Message) TableName() string {
	return MessagesTableName
}
