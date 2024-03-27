package handlers

import (
	"awesomeKonstru/backend/models"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func SaveCSVInInsumo(DSN string, route string) {
	// Conectar a la base de datos MySQL
	db, err := Connect(DSN)
	if err != nil {
		panic("Failed to conect databse")
	}

	// Migrar el esquema si no se ha hecho
	db.AutoMigrate(&models.Insumo{})

	// Abrir archivo CSV
	file, err := os.Open(route)
	if err != nil {
		panic("failed to open CSV file")
	}
	defer file.Close()

	// Crear un lector CSV
	reader := csv.NewReader(file)
	// Iterar sobre las filas del archivo CSV
	for {
		// Leer cada registro del CSV
		record, err := reader.Read()
		if err != nil {
			break
		}
		// Convertir la cantidad a float64, si hay error asignar 0
		quantity, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			fmt.Println("Error al convertir la cantidad a float64, asignando 0:", err)
			quantity = 0
		}
		fechaActualizacion, err := time.Parse("02/01/2006", record[4])
		if err != nil {
			fmt.Println("Error al asignar la fecha:", err)
			quantity = 0
		}

		insert := models.Insumo{
			ID:                 record[0],
			Descripcion:        record[1],
			Unidad:             record[2],
			PrecioBase:         quantity,
			FechaActualizacion: fechaActualizacion,
			Clasificacion:      record[5],
		}
		// Guardar el usuario en la base de datos
		db.Create(&insert)
	}
	Disconnect(db)
	fmt.Println("CSV data successfully loaded into the database")
}
