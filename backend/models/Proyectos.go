package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Proyectos struct {
	IDProyecto  string  `gorm:"primaryKey;not null;index"`
	Descripcion string  `gorm:"type:longtext;size:255"`
	Fecha       string  `gorm:"type:datetime"`
	Valor       float64 `gorm:"type:decimal(10,2)"`
	TipoObra    string  `gorm:"type:longtext;size:255"`
	UsuarioID   int
	Usuario     Usuario `gorm:"foreignKey:UsuarioID"`
}
