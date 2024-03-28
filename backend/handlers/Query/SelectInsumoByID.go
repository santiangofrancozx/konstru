package Query

import (
	"awesomeKonstru/backend/handlers"
	"awesomeKonstru/backend/models"
	"gorm.io/gorm"
)

func QueryInsumoByID(db *gorm.DB, id string) (models.Insumo, error) {
	item := models.Insumo{}
	if err := db.First(&item, id).Error; err != nil {
		return models.Insumo{}, err
	}
	return item, nil
}

func SelectInsumoByID(dsn string, id string) (models.Insumo, error) {
	db, err := handlers.Connect(dsn)
	if err != nil {
		return models.Insumo{}, err
	}
	defer handlers.Disconnect(db)

	return QueryInsumoByID(db, id)
}
