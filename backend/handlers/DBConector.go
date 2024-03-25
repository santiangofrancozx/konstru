package handlers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDatabaseIfNotExists(dsn string) error {
	// Abrir una conexión temporal sin seleccionar ninguna base de datos
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	defer Disconnect(db)

	// Ejecutar la consulta SQL para crear la base de datos si no existe
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", "KONSTRU")
	if err := db.Exec(sql).Error; err != nil {
		return err
	}

	return nil
}

func Connect(dsn string) (*gorm.DB, error) {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Crear la base de datos si no existe
	err2 := CreateDatabaseIfNotExists(dsn)
	if err2 != nil {
		return nil, err2
	}
	//Seleccionar la base de datos
	if err := db.Exec("USE KONSTRU").Error; err != nil {
		return nil, err
	}

	return db, nil
}

func Disconnect(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error al obtener el objeto DB:", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		fmt.Println("Error al cerrar la conexión:", err)
	}
}
