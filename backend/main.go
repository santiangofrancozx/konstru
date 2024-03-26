package main

import (
	"awesomeKonstru/backend/handlers"
	"awesomeKonstru/backend/models"
)

func ExecuteMigrations() {
	DSN := "root:safraroot@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	//connect to db Konstru
	db, err := handlers.Connect(DSN)
	if err != nil {
		return
	}
	//Make Migrations in Konstru
	modelsToMigrate := []interface{}{}
	modelsToMigrate = append(modelsToMigrate, &models.Usuario{})
	modelsToMigrate = append(modelsToMigrate, &models.Insumo{})
	modelsToMigrate = append(modelsToMigrate, &models.Actividad{})
	modelsToMigrate = append(modelsToMigrate, &models.ActividadInsumo{})
	errM := handlers.MakeMigrations(db, modelsToMigrate)
	if errM != nil {
		return
	}
	//Disconnect to db
	handlers.Disconnect(db)

}

func main() {
	ExecuteMigrations()
}
