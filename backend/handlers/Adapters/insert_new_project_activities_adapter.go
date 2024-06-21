package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func InsertNewProjectActivitiesAdapter(projectActivities []models.Proyectos_actividades) error {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return err
	}
	errI := Queries.InsertNewProjectActivities(db, projectActivities)
	if errI != nil {
		defer Connection_Migrates.Disconnect(db)
		return errI // Devuelve el error de inserción, no el error de conexión
	}

	defer Connection_Migrates.Disconnect(db)
	return nil
}
func DeleteProjectActivitiesAdapter(IDProyecto string, IDActivities []string) error {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return err
	}
	errI := Queries.DeleteSpecificProjectActivities(db, IDProyecto, IDActivities)
	if errI != nil {
		defer Connection_Migrates.Disconnect(db)
		return errI // Devuelve el error de inserción, no el error de conexión
	}

	defer Connection_Migrates.Disconnect(db)
	return nil
}
