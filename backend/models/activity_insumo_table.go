package models

type ActividadInsumo struct {
	ActividadID string `gorm:"primaryKey;not null"`
	InsumoID    string `gorm:"primaryKey;not null"`
	Cantidad    float64
}
