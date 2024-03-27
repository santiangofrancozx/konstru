package main

import (
	"awesomeKonstru/backend/handlers"
	"awesomeKonstru/backend/models"
)

func ExecuteMigrations(DSN string) {
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

func ImportDataFromCSVDB(DSN string) {
	CSV := "CSV-DB/insumoActividad.csv"
	CSV2 := "CSV-DB/actividades.csv"
	CSV3 := "CSV-DB/insumos.csv"
	// migra los archivos csv a la db, si ya se ejecuto una vez no es necesario mas no manda error solo no los
	//ingresa ya que se genera dulicidad de pk, estos se ejecutan solo en modo development.
	handlers.SaveCSVInInsumo(DSN, CSV3)
	handlers.SaveCSVInTableInsumoActividad(DSN, CSV)
	handlers.SaveCSVInActividad(DSN, CSV2)
}

func main() {
	DSN := "" // Define aqui tu DSN para desarrollo en local
	// strutura de DSN user:password@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local
	//cambia user por tu usuario de MySql y password por tu contrasenia de acceso a el puerto, por defecto uso 3306
	ExecuteMigrations(DSN) // ejecut las migraciones crea la db si no existe y las tablas en esta si no existen
	ImportDataFromCSVDB(DSN)
}
