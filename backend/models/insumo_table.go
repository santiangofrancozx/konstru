package models

type Insumo struct {
	ID                 uint    `gorm:"primaryKey;autoIncrement"`
	Descripcion        string  `gorm:"size:100"`
	Unidad             string  `gorm:"size:20"`
	PrecioBase         float64 `gorm:"type:decimal(10,2)"`
	FechaActualizacion string  `gorm:"type:datetime"`
	Clasificacion      string  `gorm:"size:50"`
}
