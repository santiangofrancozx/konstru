package models

type ActividadInsumo struct {
	ActividadID string `gorm:"primaryKey;not null;index"`
	InsumoID    string `gorm:"primaryKey;not null;index"`
	Cantidad    float64
}
