package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func InsertNewUser(user models.Usuario) (bool, error) {
	// Conectar a la base de datos
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return false, err
	}
	successful, err := Queries.QueryInsertNewUser(db, &user)
	if err != nil {
		return false, err
	}

	defer Connection_Migrates.Disconnect(db)

	// Llamar a la función QueryInsertNewUser

	return successful, nil
}
