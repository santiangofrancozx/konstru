package models

import "time"

type Insumo struct {
	ID                 string    `gorm:"primaryKey;not null,index"`
	Descripcion        string    `gorm:"size:100"`
	Unidad             string    `gorm:"size:20"`
	PrecioBase         float64   `gorm:"type:decimal(10,2)"`
	FechaActualizacion time.Time `gorm:"type:datetime"`
	Clasificacion      string    `gorm:"size:50;index"`
}
