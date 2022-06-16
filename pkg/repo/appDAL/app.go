package appDAL

import (
	"github.com/MahmoudMekki/ChatSystem/database"
	"github.com/MahmoudMekki/ChatSystem/pkg/models"
)

func CreateApp(app models.Application) (models.Application, error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.Application{}, err
	}
	err = dbConn.Table(models.ApplicationsTableName).Create(&app).Error
	if err != nil {
		return app, err
	}
	return app, nil
}
func UpdateApp(app models.Application) (models.Application, error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.Application{}, err
	}
	err = dbConn.Table(models.ApplicationsTableName).Where("token=?", app.Token).Updates(&app).Error
	if err != nil {
		return models.Application{}, err
	}
	return app, nil
}

func GetAppByToken(token string) (app models.Application, err error) {
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		return models.Application{}, err
	}
	err = dbConn.Table(models.ApplicationsTableName).Where("token=?", token).Find(&app).Error
	if err != nil {
		return app, err
	}
	return app, nil
}
