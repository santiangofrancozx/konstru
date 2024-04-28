package Query

import (
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
	"log"

	"gorm.io/gorm"
)

func QueryInsertNewUser(db *gorm.DB, user *models.Usuario) (bool, error) {

	if err := db.Create(user).Error; err != nil {
		log.Println("Exitoso")
		return false, err
	}

	return true, nil
}

func InsertNewUser(user models.Usuario) (bool, error) {
	// Conectar a la base de datos
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return false, err
	}
	defer Connection_Migrates.Disconnect(db)

	// Llamar a la función QueryInsertNewUser
	successful, err := QueryInsertNewUser(db, &user)
	if err != nil {
		return false, err
	}

	return successful, nil
}
