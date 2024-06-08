package Adapters

import (
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
)

func SelectApuByActivityId(id string) ([]Queries.APUs, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return nil, err
	}
	return Queries.QueryAPUByActivityID(db, id)

}
func SelectApuUserByActivityId(id string) ([]Queries.APUs, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return nil, err
	}
	return Queries.QueryAPUUserByActivityID(db, id)

}
