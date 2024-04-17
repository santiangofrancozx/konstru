package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ActividadU_InsumoU struct {
	ActividadUID      string `gorm:"type:longtext;size:255"` // Specify the length for actividad_id
	InsumoUID         string `gorm:"type:longtext;size:255"` // Specify the length for insumo_id
	Cantidad          float64
	Insumos_Usuario   Insumos_Usuario   `gorm:"foreignKey:InsumoUID"`
	Actividad_Usuario Actividad_Usuario `gorm:"foreignKey:ActividadUID"`
}
