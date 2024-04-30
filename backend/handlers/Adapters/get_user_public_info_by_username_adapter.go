package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
)

func SelectUserByUsernameInfo(email string) (Queries.UserInfo, error) {
	var userF Queries.UserInfo
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return Queries.UserInfo{}, err
	}
	defer Connection_Migrates.Disconnect(db)
	user, err := Queries.QueryUserByUsername(db, email)
	if err != nil {
		return Queries.UserInfo{}, err
	}
	userF.Email = user.Email
	userF.Nombre = user.Nombre
	userF.Apellido = user.Apellido

	return userF, err
}
