package models

type Insumos_Usuario struct {
	ID_Actividad_Usuario string `gorm:"primaryKey;not null;index"`
	Usuario_ID           string
	Descripcion          string  `gorm:"size:100;index"`
	Unidad               string  `gorm:"size:20;not null"`
	PrecioBase           float64 `gorm:"type:decimal(10,2)"`
	FechaCreacion        string  `gorm:"type:datetime"`
}
