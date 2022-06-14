package models

import "time"

const (
	ApplicationsTableName = "apps"
)

type Application struct {
	Id        int       `gorm:"column:id;primary_key" json:"-"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Token     string    `gorm:"column:token;unique" json:"token"`
	CreatedAt time.Time `gorm:"column:created_at;default: CURRENT_TIMESTAMP;not null" json:"-"`
}

func (a *Application) TableName() string {
	return ApplicationsTableName
}
