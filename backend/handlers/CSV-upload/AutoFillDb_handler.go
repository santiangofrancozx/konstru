package CSV_upload

import (
	"awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func SaveCSVInTableInsumoActividad(route string) {
	// Conectar a la base de datos MySQL
	db, err := Connection_Migrates.Connect()
	if err != nil {
		panic("Failed to conect databse")
	}

	// Migrar el esquema si no se ha hecho
	//db.AutoMigrate(&models.Actividad{})

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
		// Crear una instancia de User con los datos del CSV
		quantity, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			panic("Error al convertir la cantidad a float64")
		}
		insert := models.ActividadInsumo{
			ActividadID: record[0],
			InsumoID:    record[1],
			Cantidad:    quantity,
		}
		// Guardar el usuario en la base de datos
		db.Create(&insert)
	}
	Connection_Migrates.Disconnect(db)
	fmt.Println("CSV data successfully loaded into the database")
}
