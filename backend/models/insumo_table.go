package models

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type Insumo struct {
	ID            string  `gorm:"primaryKey;not null,index"`
	Descripcion   string  `gorm:"size:100"`
	Unidad        string  `gorm:"size:20"`
	PrecioBase    float64 `gorm:"type:decimal(10,2)"`
	Clasificacion string  `gorm:"size:50;index"`
	Base
}

func (insumo *Insumo) BeforeCreate(tx *gorm.DB) (err error) {
	if insumo.ID == "" {
		insumo.ID = uuid.New().String()
	}
	return
}
