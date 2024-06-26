package Migrates

import (
	CSV_upload "awesomeKonstru/backend/handlers/CSV-upload"
	Connection_Migrates "awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
	"fmt"
	"gorm.io/gorm"
)

type Migration interface {
	Migrate() error
}

func MakeMigrations(models []interface{}) error {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return nil
	}
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("error al realizar las migraciones: %v", err)
	}
	consultaSQL := "CREATE UNIQUE INDEX unique_actividad_insumo ON actividad_insumos (actividad_id, insumo_id);"
	consultaSQL3 := "CREATE UNIQUE INDEX unique_proyectos_actividades ON proyectos_actividades (id_proyecto, id_actividad);"
	consultaSQL2 := "CREATE UNIQUE INDEX unique_actividad_insumo_user ON actividadu_insumo_us (actividad_uid, insumo_uid);"
	// Ejecutar la consulta SQL
	if err := db.Exec(consultaSQL).Error; err != nil {
		return err
	}
	if err := db.Exec(consultaSQL2).Error; err != nil {
		return err
	}
	if err := db.Exec(consultaSQL3).Error; err != nil {
		return err
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
	modelsToMigrate = append(modelsToMigrate, &models.Actividad_Usuario{})
	modelsToMigrate = append(modelsToMigrate, &models.ActividadU_InsumoU{})
	modelsToMigrate = append(modelsToMigrate, &models.Proyectos{})
	modelsToMigrate = append(modelsToMigrate, &models.Proyectos_actividades{})

	return modelsToMigrate

}
func CreateAdminUser() error {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return err
	}
	defer Connection_Migrates.Disconnect(db)

	// Verificar si el usuario admin ya existe
	var existingUser models.Usuario
	if err := db.Where("username = ?", "admin").First(&existingUser).Error; err == nil {
		// El usuario ya existe, no hacer nada
		return nil
	} else if err != gorm.ErrRecordNotFound {
		// Otro error ocurrió
		return err
	}

	// Crear el usuario admin
	adminUser := models.Usuario{
		Nombre:   "admin",
		Apellido: "admin",
		Password: "admin_password_hash", // Asegúrate de usar un hash de contraseña seguro
		Email:    "admin@example.com",
	}

	if err := db.Create(&adminUser).Error; err != nil {
		return fmt.Errorf("error al crear usuario admin: %v", err)
	}

	return nil
}

func ImportDataFromCSVDB() {
	// migra los archivos csv a la db, si ya se ejecuto una vez no es necesario mas no manda error solo no los
	//ingresa ya que se genera dulicidad de pk, estos se ejecutan solo en modo development.
	CSV3 := "CSV-DB/insumos.csv"
	CSV_upload.SaveCSVInInsumo(CSV3)
	CSV2 := "CSV-DB/actividades.csv"
	CSV_upload.SaveCSVInActividad(CSV2)
	CSV := "CSV-DB/insumoActividad.csv"
	CSV_upload.SaveCSVInTableInsumoActividad(CSV)
}
