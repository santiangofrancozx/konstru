package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ActividadInsumo struct {
	ActividadID string `gorm:"type:longtext;size:255"` // Specify the length for actividad_id
	InsumoID    string `gorm:"type:longtext;size:255"` // Specify the length for insumo_id
	Cantidad    float64
	Insumo      Insumo    `gorm:"foreignKey:InsumoID"`
	Actividad   Actividad `gorm:"foreignKey:ActividadID"`
}
