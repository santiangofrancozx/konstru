package handlers

import (
	"awesomeKonstru/backend/models"
	"fmt"
	"gorm.io/gorm"
)

func MakeMigrations(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Usuario{}, &models.Insumo{}, &models.Actividad{}, &models.ActividadInsumo{}); err != nil {
		return fmt.Errorf("error al realizar las migraciones: %v", err)
	}
	return nil
}
