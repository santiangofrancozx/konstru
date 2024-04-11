package models

type ActividadInsumo struct {
	ActividadID string `gorm:"foreignKey:ActividadID;references:Actividades.ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	InsumoID    string `gorm:"foreignKey:InsumoID;references:Insumos.ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Cantidad    float64
}
