package Query

import (
	"awesomeKonstru/backend/handlers/Connection-Migrates"
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

func SelectUserByUsername(email string) (models.Usuario, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return models.Usuario{}, err
	}
	user, err := QueryUserByUsername(db, email)
	defer Connection_Migrates.Disconnect(db)

	return user, err
}

func SelectUserByUsernameInfo(email string) (UserInfo, error) {
	var userF UserInfo
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return UserInfo{}, err
	}
	defer Connection_Migrates.Disconnect(db)
	user, err := QueryUserByUsername(db, email)
	if err != nil {
		return UserInfo{}, err
	}
	userF.Email = user.Email
	userF.Nombre = user.Nombre
	userF.Apellido = user.Apellido
	userF.FechaCreacion = user.FechaCreacion

	return userF, err
}
