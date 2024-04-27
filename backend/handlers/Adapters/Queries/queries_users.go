package Queries

import (
	"awesomeKonstru/backend/models"
	"gorm.io/gorm"
)

type UserInfo struct {
	Nombre        string
	Apellido      string
	Email         string
	FechaCreacion string
}

func QueryUserByUsername(db *gorm.DB, email string) (models.Usuario, error) {
	user := models.Usuario{}
	if err := db.First(&user, "email = ?", email).Error; err != nil {
		return models.Usuario{}, err
	}
	return user, nil
}
