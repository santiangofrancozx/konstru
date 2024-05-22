package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	"awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func InsertIntoInsumoAdapter(insumo models.Insumo) error {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return err
	}

	errI := Queries.InsertInsumo(db, insumo)
	if errI != nil {
		defer Connection_Migrates.Disconnect(db)
		return errI // Devuelve el error de inserción, no el error de conexión
	}

	defer Connection_Migrates.Disconnect(db)
	return nil
}
