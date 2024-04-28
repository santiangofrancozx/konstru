package Queries

import (
	"awesomeKonstru/backend/models"

	"gorm.io/gorm"
)

func InsertInToInsumosUsuario(db *gorm.DB, insumoUsuario models.Insumos_Usuario) (bool, error) {
	if err := db.Create(&insumoUsuario).Error; err != nil {
		return false, err
	}
	return true, nil
}
