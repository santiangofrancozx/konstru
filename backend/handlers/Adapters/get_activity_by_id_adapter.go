package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func SelectActivityByID(id string) (models.Actividad, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return models.Actividad{}, err
	}
	defer Connection_Migrates.Disconnect(db)

	return Queries.QueryActivityByID(db, id)
}
func SelectActivityByUserID(id string) ([]models.Actividad_Usuario, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return []models.Actividad_Usuario{}, err
	}
	defer Connection_Migrates.Disconnect(db)

	return Queries.QueryActivityByUserID(db, id)
}

func SelectActivityByNameWithUserID(name string, id string) ([]models.Actividad_Usuario, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return []models.Actividad_Usuario{}, err
	}
	defer Connection_Migrates.Disconnect(db)

	return Queries.QueryActividadUserByNombre(db, name, id)
}

func SelectActivityUserById(idU string, id string) (models.Actividad_Usuario, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return models.Actividad_Usuario{}, err
	}
	defer Connection_Migrates.Disconnect(db)

	return Queries.QueryActividadUserById(db, idU, id)
}
