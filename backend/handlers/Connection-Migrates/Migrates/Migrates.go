package Migrates

import (
	"awesomeKonstru/backend/handlers/CSV-upload"
	"awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
	"fmt"
	"gorm.io/gorm"
)

type Migration interface {
	Migrate() error
}

func MakeMigrations(db *gorm.DB, models []interface{}) error {
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("error al realizar las migraciones: %v", err)
	}
	return nil
}
func ExecuteMigrations(DSN string) {
	//connect to db Konstru
	db, err := Connection_Migrates.Connect(DSN)
	if err != nil {
		return
	}
	//Make Migrations in Konstru
	modelsToMigrate := []interface{}{}
	modelsToMigrate = append(modelsToMigrate, &models.Usuario{})
	modelsToMigrate = append(modelsToMigrate, &models.Insumo{})
	modelsToMigrate = append(modelsToMigrate, &models.Actividad{})
	modelsToMigrate = append(modelsToMigrate, &models.ActividadInsumo{})

	errM := MakeMigrations(db, modelsToMigrate)
	if errM != nil {
		return
	}
	//Disconnect to db
	Connection_Migrates.Disconnect(db)

}

func ImportDataFromCSVDB(DSN string) {
	CSV := "CSV-DB/insumoActividad.csv"
	CSV2 := "CSV-DB/actividades.csv"
	CSV3 := "CSV-DB/insumos.csv"
	// migra los archivos csv a la db, si ya se ejecuto una vez no es necesario mas no manda error solo no los
	//ingresa ya que se genera dulicidad de pk, estos se ejecutan solo en modo development.
	CSV_upload.SaveCSVInInsumo(DSN, CSV3)
	CSV_upload.SaveCSVInTableInsumoActividad(DSN, CSV)
	CSV_upload.SaveCSVInActividad(DSN, CSV2)
}
