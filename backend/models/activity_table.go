package models

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type Actividad struct {
	ID               string            `gorm:"primaryKey;not null;index"`
	OriginalID       string            `gorm:"size:255"`
	Descripcion      string            `gorm:"size:100;index"`
	Unidad           string            `gorm:"size:20;not null"`
	PrecioBase       float64           `gorm:"type:decimal(10,2)"`
	CreatedBy        *string           `gorm:"size:255"`
	Usuario          Usuario           `gorm:"foreignKey:CreatedBy"`
	ActividadInsumos []ActividadInsumo `gorm:"foreignKey:ActividadID;constraint:OnDelete:CASCADE;"`
	Base
}

func (b *Actividad) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}
