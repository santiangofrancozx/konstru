package models

func MigrateDB() {
	DB.AutoMigrate(Usuario{}, &Insumo{}, &Actividad{}, &ActividadInsumo{})
}
