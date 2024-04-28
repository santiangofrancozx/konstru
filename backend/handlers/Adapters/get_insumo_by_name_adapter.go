package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
)

func SelectInsumoByNombre(name string) ([]models.Insumo, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return nil, err
	}
	defer Connection_Migrates.Disconnect(db)

	return Queries.QueryInsumoByNombre(db, name)
}
