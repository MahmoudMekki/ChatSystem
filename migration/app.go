package migration

import (
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func listApp() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "create-applications-table-and-token-index",
			Migrate: func(db *gorm.DB) error {
				return db.AutoMigrate(&models.Application{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable(models.ApplicationsTableName)
			},
		},
	}
}
