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
