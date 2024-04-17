package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Insumos_Usuario struct {
	ID_Insumo_Usuario string `gorm:"primaryKey;type:varchar(255)"`
	Usuario_ID        int
	Descripcion       string  `gorm:"size:100;index"`
	Unidad            string  `gorm:"size:20;not null"`
	PrecioBase        float64 `gorm:"type:decimal(10,2)"`
	FechaCreacion     string  `gorm:"type:datetime"`
	Usuario           Usuario `gorm:"foreignKey:Usuario_ID"`
}
