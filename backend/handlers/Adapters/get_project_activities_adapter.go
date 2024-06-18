package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func SelectAllProjectsActivities(idp string) ([]models.Proyectos_actividades, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return nil, err
	}
	defer Connection_Migrates.Disconnect(db)

	return Queries.GetAllProjectActivities(db, idp)
}
