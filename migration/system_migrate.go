package migration

import (
	"github.com/MahmoudMekki/ChatSystem/database"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/rs/zerolog/log"
)

func RunMigration() {
	log.Info().Msg("Migration is started")
	var chatSystem []*gormigrate.Migration
	chatSystem = append(chatSystem, listApp()...)
	chatSystem = append(chatSystem, listChat()...)
	chatSystem = append(chatSystem, listMessages()...)
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	migration := gormigrate.New(dbConn, gormigrate.DefaultOptions, chatSystem)
	if err = migration.Migrate(); err != nil {
		log.Fatal().Msg("Couldn't run the migration")
	}
	log.Info().Msg("Migration is done successfully")
}
