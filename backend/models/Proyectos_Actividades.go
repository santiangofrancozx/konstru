package models

type Proyectos_actividades struct {
	ID_proyecto  string `gorm:"type:longtext;size:255"`
	ID_actividad string `gorm:"type:longtext;size:255"`
	Cantidad     float64
	Proyectos    Proyectos `gorm:"foreignKey:ID_proyecto;constraint:OnDelete:CASCADE;"`
	Actividad    Actividad `gorm:"foreignKey:ID_actividad;constraint:OnDelete:CASCADE;"`
}
