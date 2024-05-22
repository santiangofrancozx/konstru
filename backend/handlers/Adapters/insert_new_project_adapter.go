package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	"awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func InsertNewProjectAdapter(project models.Proyectos) error {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return err
	}

	errI := Queries.InsertNewProject(db, &project)
	if errI != nil {
		defer Connection_Migrates.Disconnect(db)
		return errI // Devuelve el error de inserción, no el error de conexión
	}

	defer Connection_Migrates.Disconnect(db)
	return nil
}
