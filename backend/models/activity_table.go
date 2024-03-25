package models

type Actividad struct {
	ID            uint    `gorm:"primaryKey;autoIncrement"`
	Descripcion   string  `gorm:"size:100"`
	Unidad        string  `gorm:"size:20"`
	ValorTotal    float64 `gorm:"type:decimal(10,2)"`
	FechaCreacion string  `gorm:"type:datetime"`
}
