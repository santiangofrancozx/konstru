package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Actividad_Usuario struct {
	ID_Actividad_Usuario int     `gorm:"primaryKey;not null;autoIncrement;index"`
	Usuario_ID           string  `gorm:"size:100;index"`
	Descripcion          string  `gorm:"size:100;index"`
	Unidad               string  `gorm:"size:20;not null"`
	PrecioBase           float64 `gorm:"type:decimal(10,2)"`
	Base
	Usuario Usuario `gorm:"foreignKey:Usuario_ID"`
}
