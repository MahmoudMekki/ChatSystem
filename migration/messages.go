package migration

import (
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func listMessages() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "create-messages-table",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(&models.Message{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable(models.MessagesTableName)
			},
		},
	}
}
