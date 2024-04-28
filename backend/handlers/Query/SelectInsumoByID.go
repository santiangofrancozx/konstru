package Query

import (
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"

	"gorm.io/gorm"
)

func QueryInsumoByID(db *gorm.DB, id string) (models.Insumo, error) {
	item := models.Insumo{}
	if err := db.First(&item, "id = ?", id).Error; err != nil {
		return models.Insumo{}, err
	}
	return item, nil
}

func QueryInsumoByNombre(db *gorm.DB, nombre string) ([]models.Insumo, error) {
	var items []models.Insumo
	if err := db.Where("descripcion LIKE ?", "%"+nombre+"%").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func SelectInsumoByID(id string) (models.Insumo, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return models.Insumo{}, err
	}
	defer Connection_Migrates.Disconnect(db)

	return QueryInsumoByID(db, id)
}

func SelectInsumoByNombre(name string) ([]models.Insumo, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return nil, err
	}
	defer Connection_Migrates.Disconnect(db)

	return QueryInsumoByNombre(db, name)
}
