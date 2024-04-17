package Migrates

import (
	CSV_upload "awesomeKonstru/backend/handlers/CSV-upload"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
	"fmt"
)

type Migration interface {
	Migrate() error
}

func MakeMigrations(DSN string, models []interface{}) error {
	db, err := Connection_Migrates.Connect(DSN)
	if err != nil {
		return nil
	}
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("error al realizar las migraciones: %v", err)
	}
	Connection_Migrates.Disconnect(db)
	return nil
}

func ExecuteMigrations() []interface{} {
	//connect to db Konstru
	//Make Migrations in Konstru
	modelsToMigrate := []interface{}{}
	modelsToMigrate = append(modelsToMigrate, &models.Usuario{})
	modelsToMigrate = append(modelsToMigrate, &models.Insumo{})
	modelsToMigrate = append(modelsToMigrate, &models.Actividad{})
	modelsToMigrate = append(modelsToMigrate, &models.ActividadInsumo{})
	modelsToMigrate = append(modelsToMigrate, &models.Insumos_Usuario{})
	modelsToMigrate = append(modelsToMigrate, &models.Actividad_Usuario{})
	modelsToMigrate = append(modelsToMigrate, &models.ActividadU_InsumoU{})
	modelsToMigrate = append(modelsToMigrate, &models.Proyectos{})
	modelsToMigrate = append(modelsToMigrate, &models.Proyectos_actividades{})

	return modelsToMigrate

}

func ImportDataFromCSVDB(DSN string) {
	// migra los archivos csv a la db, si ya se ejecuto una vez no es necesario mas no manda error solo no los
	//ingresa ya que se genera dulicidad de pk, estos se ejecutan solo en modo development.
	CSV3 := "CSV-DB/insumos.csv"
	CSV_upload.SaveCSVInInsumo(DSN, CSV3)
	CSV2 := "CSV-DB/actividades.csv"
	CSV_upload.SaveCSVInActividad(DSN, CSV2)
	CSV := "CSV-DB/insumoActividad.csv"
	CSV_upload.SaveCSVInTableInsumoActividad(DSN, CSV)
}
