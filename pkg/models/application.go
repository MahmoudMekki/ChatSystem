package models

const (
	ApplicationsTableName = "apps"
)

type Application struct {
	Id    int    `gorm:"column:id;primary_key" json:"-"`
	Name  string `gorm:"column:name;not null" json:"name"`
	Token string `gorm:"column:token;unique" json:"token"`
	Chats []Chat `json:"chats,omitempty"`
}

func (a *Application) TableName() string {
	return ApplicationsTableName
}
