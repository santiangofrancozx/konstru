package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Actividad_Usuario struct {
	ID_Actividad_Usuario string `gorm:"primaryKey;not null;index"`
	Usuario_ID           int
	Descripcion          string  `gorm:"size:100;index"`
	Unidad               string  `gorm:"size:20;not null"`
	PrecioBase           float64 `gorm:"type:decimal(10,2)"`
	FechaCreacion        string  `gorm:"type:datetime"`
	Usuario              Usuario `gorm:"foreignKey:Usuario_ID"`
}
