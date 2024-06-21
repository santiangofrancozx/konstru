package Queries

import (
	"awesomeKonstru/backend/models"
	"gorm.io/gorm"
	"time"
)

type APUs struct {
	ID                 string
	Descripcion        string
	Unidad             string
	PrecioBase         float64
	FechaActualizacion time.Time
	Clasificacion      string
	Cantidad           float64
}

func QueryAPUByActivityID(db *gorm.DB, id string) ([]APUs, error) {
	var insumos []APUs
	if err := db.
		Model(&models.ActividadInsumo{}).
		Select("insumos.*, actividad_insumos.cantidad").
		Joins("JOIN insumos ON actividad_insumos.insumo_id = insumos.id").
		Where("actividad_insumos.actividad_id = ?", id).
		Find(&insumos).Error; err != nil {
		return nil, err
	}
	return insumos, nil
}
func QueryAPUUserByActivityID(db *gorm.DB, id string) ([]APUs, error) {
	var insumos []APUs
	if err := db.
		Model(&models.ActividadU_InsumoU{}).
		Select("insumos.*, actividadu_insumo_us.cantidad").
		Joins("JOIN insumos ON actividadu_insumo_us.insumo_uid = insumos.id").
		Where("actividadu_insumo_us.actividad_uid = ?", id).
		Find(&insumos).Error; err != nil {
		return nil, err
	}
	return insumos, nil
}

func QueryInsertNewAPUs(db *gorm.DB, apu []models.ActividadInsumo) error {
	if err := db.Create(&apu).Error; err != nil {
		return err
	}
	return nil
}
func QueryInsertNewAPUsUser(db *gorm.DB, apu []models.ActividadU_InsumoU) error {
	if err := db.Create(&apu).Error; err != nil {
		return err
	}
	return nil
}
