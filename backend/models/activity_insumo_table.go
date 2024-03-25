package models

type ActividadInsumo struct {
	ActividadID uint `gorm:"primaryKey"`
	InsumoID    uint `gorm:"primaryKey"`
	Cantidad    int
}
