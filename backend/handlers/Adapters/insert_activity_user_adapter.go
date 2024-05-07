package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func InsertIntoActivityUserAdapter(activity models.Actividad_Usuario) error {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return err
	}

	errI := Queries.InsertInToActivitiesUser(db, activity)
	if errI != nil {
		defer Connection_Migrates.Disconnect(db)
		return errI // Devuelve el error de inserción, no el error de conexión
	}

	defer Connection_Migrates.Disconnect(db)
	return nil
}
