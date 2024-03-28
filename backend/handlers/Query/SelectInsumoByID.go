package Query

import (
	"awesomeKonstru/backend/handlers/Connection-Migrates"
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
	db, err := Connection_Migrates.Connect(dsn)
	if err != nil {
		return models.Insumo{}, err
	}
	defer Connection_Migrates.Disconnect(db)

	return QueryInsumoByID(db, id)
}
