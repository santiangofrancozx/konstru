package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ActividadInsumo struct {
	ActividadID         string `gorm:"type:longtext;size:255"`
	OriginalActividadID string `gorm:"type:longtext;size:255"`
	InsumoID            string `gorm:"type:longtext;size:255"`
	Cantidad            float64
	Insumo              Insumo    `gorm:"foreignKey:InsumoID"`
	Actividad           Actividad `gorm:"foreignKey:ActividadID"`
}
