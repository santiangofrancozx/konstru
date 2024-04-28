package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	"awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func InserteIntoInsumoUsuarioAdapter(user models.Insumos_Usuario) (bool, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return false, err
	}

	response, err := Queries.InsertInToInsumosUsuario(db, user)
	if err != nil {
		defer Connection_Migrates.Disconnect(db)
		return false, err
	}

	defer Connection_Migrates.Disconnect(db)
	return response, err

}
