package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Usuario representa la estructura del modelo Usuario

type Usuario struct {
	ID string `gorm:"primary_key"`
	Base
	Nombre   string `gorm:"size:50"`
	Apellido string `gorm:"size:50"`
	Email    string `gorm:"size:100;unique;index"`
	Password string `gorm:"size:100"` // Corregido a Password con P mayúscula
}

func (b *Usuario) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
