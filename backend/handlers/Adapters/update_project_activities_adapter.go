package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func UpdateProjectActivitiesAdapter(PAmodels []models.Proyectos_actividades, Idproject string) error {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return err
	}

	succesful := Queries.UpdateProjectActivities(db, PAmodels, Idproject)
	if succesful != nil {
		defer Connection_Migrates.Disconnect(db)
		return succesful
	}
	defer Connection_Migrates.Disconnect(db)
	return nil
}
