package handlers

import (
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
	db, err := Connect(DSN)
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
	Disconnect(db)

}

func ImportDataFromCSVDB(DSN string) {
	CSV := "CSV-DB/insumoActividad.csv"
	CSV2 := "CSV-DB/actividades.csv"
	CSV3 := "CSV-DB/insumos.csv"
	// migra los archivos csv a la db, si ya se ejecuto una vez no es necesario mas no manda error solo no los
	//ingresa ya que se genera dulicidad de pk, estos se ejecutan solo en modo development.
	SaveCSVInInsumo(DSN, CSV3)
	SaveCSVInTableInsumoActividad(DSN, CSV)
	SaveCSVInActividad(DSN, CSV2)
}
