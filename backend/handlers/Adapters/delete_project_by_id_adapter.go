package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
)

func DeleteProjectByIdAdapter(id string) error {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return err
	}
	defer Connection_Migrates.Disconnect(db)
	return Queries.DeleteProjectByID(db, id)
}
