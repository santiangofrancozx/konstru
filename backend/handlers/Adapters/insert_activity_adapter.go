package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func InsertIntoActivityAdapter(activity models.Actividad) (error, string) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return err, "nil"
	}

	errI, lastId := Queries.InsertInToActivities(db, activity)
	if errI != nil {
		defer Connection_Migrates.Disconnect(db)
		return errI, "" // Devuelve el error de inserción, no el error de conexión
	}

	defer Connection_Migrates.Disconnect(db)
	return nil, lastId
}
