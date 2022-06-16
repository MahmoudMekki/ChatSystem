package migration

import (
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func listChat() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "create-chat-table",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(&models.Chat{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable(models.ChatTableName)
			},
		},
	}
}
