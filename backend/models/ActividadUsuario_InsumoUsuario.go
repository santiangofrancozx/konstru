package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ActividadU_InsumoU struct {
	ActividadUID      int               `gorm:"not null;index"` // Specify the length for actividad_id
	InsumoUID         string            `gorm:"not null;index"` // Specify the length for insumo_id
	Cantidad          float64           `gorm:"not null"`
	Insumos           Insumo            `gorm:"foreignKey:InsumoUID"`
	Actividad_Usuario Actividad_Usuario `gorm:"foreignKey:ActividadUID"`
}
