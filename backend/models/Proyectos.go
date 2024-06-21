package models

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type Proyectos struct {
	IDProyecto  string  `gorm:"primaryKey;not null;index"`
	Name        string  `gorm:"type:longtext;size:255"`
	Descripcion string  `gorm:"type:longtext;size:255"`
	Valor       float64 `gorm:"type:decimal(10,2)"`
	TipoObra    string  `gorm:"type:longtext;size:255"`
	State       bool    `gorm:"default:true"`
	Base
	UsuarioID   string                  `gorm:"type:longtext;size:255"`
	Usuario     Usuario                 `gorm:"foreignKey:UsuarioID"`
	Actividades []Proyectos_actividades `gorm:"foreignKey:ID_proyecto;constraint:OnDelete:CASCADE;"`
}

func (b *Proyectos) BeforeCreate(tx *gorm.DB) (err error) {
	b.IDProyecto = uuid.New().String()
	return
}
