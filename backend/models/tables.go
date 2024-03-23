package models

// Usuario representa la estructura del modelo Usuario
type Usuario struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Nombre        string `gorm:"size:50"`
	Apellido      string `gorm:"size:50"`
	Email         string `gorm:"size:100"`
	Contraseña    string `gorm:"size:100"`
	FechaCreacion string `gorm:"type:datetime"`
}

// Insumo representa la estructura del modelo Insumo
type Insumo struct {
	ID                 uint    `gorm:"primaryKey;autoIncrement"`
	Descripcion        string  `gorm:"size:100"`
	Unidad             string  `gorm:"size:20"`
	PrecioBase         float64 `gorm:"type:decimal(10,2)"`
	FechaActualizacion string  `gorm:"type:datetime"`
	Clasificacion      string  `gorm:"size:50"`
}

// Actividad representa la estructura del modelo Actividad
type Actividad struct {
	ID            uint    `gorm:"primaryKey;autoIncrement"`
	Descripcion   string  `gorm:"size:100"`
	Unidad        string  `gorm:"size:20"`
	ValorTotal    float64 `gorm:"type:decimal(10,2)"`
	FechaCreacion string  `gorm:"type:datetime"`
}

// ActividadInsumo representa la estructura del modelo ActividadInsumo
type ActividadInsumo struct {
	ActividadID uint `gorm:"primaryKey"`
	InsumoID    uint `gorm:"primaryKey"`
	Cantidad    int
}
