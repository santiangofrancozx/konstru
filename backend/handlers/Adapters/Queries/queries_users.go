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

func QueryInsertNewUser(db *gorm.DB, user *models.Usuario) (bool, error) {
	if err := db.Create(user).Error; err != nil {
		return false, err
	}

	return true, nil
}
