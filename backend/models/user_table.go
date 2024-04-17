package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Usuario representa la estructura del modelo Usuario
type Usuario struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Nombre        string `gorm:"size:50"`
	Apellido      string `gorm:"size:50"`
	Email         string `gorm:"size:100;index"`
	password      string `gorm:"size:100"`
	FechaCreacion string `gorm:"type:datetime"`
}
