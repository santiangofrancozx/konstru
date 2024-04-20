package Query

import (
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
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

func SelectApuByActivityId(id string) ([]APUs, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return nil, err
	}
	return QueryAPUByActivityID(db, id)
}
