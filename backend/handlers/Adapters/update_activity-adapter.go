package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
	"fmt"
)

func UpdateActivityAdapter(activity models.Actividad) error {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return err
	}
	defer Connection_Migrates.Disconnect(db)

	// Log para verificar el ID antes de pasar a la consulta
	fmt.Println("Updating actividad with ID:", activity.ID)

	errU := Queries.UpdateActividad(db, activity)
	if errU != nil {
		return errU
	}

	return nil
}

func UpdateActivityInsumosAdapter(id string, insumos []models.ActividadInsumo) error {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return err
	}
	activityToUpdate, errA := Queries.QueryActivityByID(db, id)
	if errA != nil {
		return errA
	}

	errD := Queries.DeleteAPUByID(db, activityToUpdate.ID)
	if errD != nil {
		return errD
	}

	errIU := Queries.QueryInsertNewAPUs(db, insumos)
	if errIU != nil {
		return errIU
	}
	return nil
}
