package main

import (
	"awesomeKonstru/backend/handlers"
)

func main() {
	DSN := "root:safraroot@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	//connect to db Konstru
	db, err := handlers.Connect(DSN)
	if err != nil {
		return
	}
	//Make Migrations in Konstru
	errM := handlers.MakeMigrations(db)
	if errM != nil {
		return
	}
	//Disconnect to db
	handlers.Disconnect(db)
}
