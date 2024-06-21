package CSV_upload

import (
	"awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func SaveCSVInActividad(route string) {
	// Conectar a la base de datos MySQL
	db, err := Connection_Migrates.Connect()
	if err != nil {
		panic("Failed to conect databse")
	}

	// Migrar el esquema si no se ha hecho
	//db.AutoMigrate(&models.ActividadInsumo{})

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
		insert := models.Actividad{
			OriginalID:  record[0],
			Descripcion: record[1],
			Unidad:      record[2],
			PrecioBase:  quantity,
		}
		// Guardar el usuario en la base de datos
		db.Create(&insert)
	}
	Connection_Migrates.Disconnect(db)
	fmt.Println("CSV data successfully loaded into the database")
}
