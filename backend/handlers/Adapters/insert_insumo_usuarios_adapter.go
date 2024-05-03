package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func InserteIntoInsumoUsuarioAdapter(user models.Insumos_Usuario) error {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return err
	}

	errI := Queries.InsertInToInsumosUsuario(db, user)
	if errI != nil {
		defer Connection_Migrates.Disconnect(db)
		return errI // Devuelve el error de inserción, no el error de conexión
	}

	defer Connection_Migrates.Disconnect(db)
	return nil
}
