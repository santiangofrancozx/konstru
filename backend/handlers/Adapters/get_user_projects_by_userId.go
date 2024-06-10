package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	"awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func GetUserProjectsByUSerIdAdapter(id string) ([]models.Proyectos, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return []models.Proyectos{}, err
	}
	projects, err := Queries.GetProjectsByUserID(db, id)
	defer Connection_Migrates.Disconnect(db)

	return projects, err
}

func GetProjectByIdAdapter(id string) (*models.Proyectos, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return &models.Proyectos{}, err
	}
	project, err := Queries.GetProjectByID(db, id)
	defer Connection_Migrates.Disconnect(db)

	return project, err
}
