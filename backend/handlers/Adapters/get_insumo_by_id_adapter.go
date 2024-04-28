package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func SelectInsumoByID(id string) (models.Insumo, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return models.Insumo{}, err
	}
	defer Connection_Migrates.Disconnect(db)

	return Queries.QueryInsumoByID(db, id)
}
