package database

import (
	"fmt"
	"github.com/MahmoudMekki/ChatSystem/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var dbConn *gorm.DB
var db dataBase

type dataBase struct {
	host     string
	port     string
	name     string
	user     string
	password string
	debug    bool
	dialect  string
}

func init() {
	db = dataBase{
		host:     config.GetEnvVar("DB_HOST"),
		port:     config.GetEnvVar("DB_PORT"),
		name:     config.GetEnvVar("DB_NAME"),
		user:     config.GetEnvVar("DB_USER"),
		password: config.GetEnvVar("DB_PASSWORD"),
	}
}

func (d *dataBase) dsn() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		d.user,
		d.password,
		d.host,
		d.port,
		d.name,
	)
	return dsn
}

func CreateDBConnection() error {
	if dbConn != nil {
		CloseDBConnection(dbConn)
	}
	dsn := db.dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	dbConn = db
	return err
}
func CloseDBConnection(conn *gorm.DB) {
	sqlDB, err := conn.DB()
	if err != nil {
		log.Err(err).Msg("Error occurred while closing a DB connection")
	}
	defer sqlDB.Close()
}
func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn, err
	}
	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}
	return dbConn, nil
}
