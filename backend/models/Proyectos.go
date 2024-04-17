package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Proyectos struct {
	id_proyecto string  `gorm:"primaryKey;not null;index"`
	descripcion string  `gorm:"type:longtext;size:255"`
	fecha       string  `gorm:"type:datetime"`
	valor       float64 `gorm:"type:decimal(10,2)"`
	tipo_obra   string  `gorm:"type:longtext;size:255"`
	usuario_id  string  `gorm:"type:longtext;size:255"`
	Usuario     Usuario `gorm:"foreignKey:usuario_id"`
}
