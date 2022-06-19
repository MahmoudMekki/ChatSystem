package migration

import (
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func listApp() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "create-applications-table-and-token-unique-index",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(&models.Application{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable(models.ApplicationsTableName)
			},
		},
		{
			ID: "add-chat-counts-column",
			Migrate: func(db *gorm.DB) error {
				sql := `alter table apps ADD COLUMN chat_count bigint DEFAULT 0`
				err := db.Exec(sql).Error
				return err
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropColumn(models.Application{}, "chat_count")
			},
		},
	}
}
