package Query

import (
	"awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
	"gorm.io/gorm"
)

func QueryActivityByID(db *gorm.DB, id string) (models.Actividad, error) {
	item := models.Actividad{}
	if err := db.First(&item, id).Error; err != nil {
		return models.Actividad{}, err
	}
	return item, nil
}

func SelectActivityByID(dsn string, id string) (models.Actividad, error) {
	db, err := Connection_Migrates.Connect(dsn)
	if err != nil {
		return models.Actividad{}, err
	}
	defer Connection_Migrates.Disconnect(db)

	return QueryActivityByID(db, id)
}
