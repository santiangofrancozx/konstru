package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func SelectUserByUsername(email string) (models.Usuario, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return models.Usuario{}, err
	}
	user, err := Queries.QueryUserByUsername(db, email)
	defer Connection_Migrates.Disconnect(db)

	return user, err
}
