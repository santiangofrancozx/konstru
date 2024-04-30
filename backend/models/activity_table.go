package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Actividad struct {
	ID          string  `gorm:"primaryKey;not null;index"`
	Descripcion string  `gorm:"size:100;index"`
	Unidad      string  `gorm:"size:20;not null"`
	PrecioBase  float64 `gorm:"type:decimal(10,2)"`
	Base
}
