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
		panic("Failed to connect to database")
	}

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
		quantity, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			fmt.Println("Error al convertir la cantidad a float64, asignando 0:", err)
			quantity = 0
		}

		// Buscar la actividad por el ID original
		var actividad models.Actividad
		db.Where("original_id = ?", record[0]).First(&actividad)

		insert := models.ActividadInsumo{
			ActividadID:         actividad.ID, // Asigna el nuevo UUID basado en el original_id
			OriginalActividadID: record[0],    // Guarda el ID original del CSV
			InsumoID:            record[1],
			Cantidad:            quantity,
		}
		// Guardar la relación actividad-insumo en la base de datos
		db.Create(&insert)
	}
	Connection_Migrates.Disconnect(db)
	fmt.Println("CSV data successfully loaded into the database")
}
