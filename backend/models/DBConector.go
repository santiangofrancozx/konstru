package models

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:safraroot@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"

	// Intenta conectar a la base de datos
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Verifica si la base de datos KONSTRU existe, si no, la crea
	if err := EnsureDatabaseExists(db); err != nil {
		return nil, err
	}

	// Conecta a la base de datos KONSTRU
	dsn = "root:safraroot@tcp(127.0.0.1:3306)/KONSTRU?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Asigna la conexión a la base de datos a la variable global DB
	DB = db

	fmt.Println("Conexión a la base de datos establecida correctamente.")
	return db, nil
}

// Función auxiliar para verificar y crear la base de datos KONSTRU si no existe
func EnsureDatabaseExists(db *gorm.DB) error {
	var result sql.NullString
	err := db.Raw("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = 'KONSTRU'").Scan(&result).Error
	if err != nil {
		return err
	}

	if !result.Valid || result.String == "" {
		// La base de datos KONSTRU no existe, se intenta crear
		if err := db.Exec("CREATE DATABASE IF NOT EXISTS KONSTRU").Error; err != nil {
			return err
		}
		fmt.Println("Base de datos KONSTRU creada exitosamente.")
	}

	return nil
}
