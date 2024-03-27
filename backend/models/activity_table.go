package models

type Actividad struct {
	ID            string  `gorm:"primaryKey;not null"`
	Descripcion   string  `gorm:"size:100"`
	Unidad        string  `gorm:"size:20;not null"`
	ValorTotal    float64 `gorm:"type:decimal(10,2)"`
	FechaCreacion string  `gorm:"type:datetime"`
}
