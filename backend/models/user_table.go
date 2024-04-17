package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Usuario representa la estructura del modelo Usuario

type Usuario struct {
	ID            int    `gorm:"primaryKey;autoIncrement"`
	Nombre        string `gorm:"size:50"`
	Apellido      string `gorm:"size:50"`
	Email         string `gorm:"size:100;index"`
	Password      string `gorm:"size:100"` // Corregido a Password con P mayúscula
	FechaCreacion string `gorm:"type:datetime"`
}
