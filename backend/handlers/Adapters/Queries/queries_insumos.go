package Queries

import (
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
