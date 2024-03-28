package main

import (
	"awesomeKonstru/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//DSN := "" // Define aqui tu DSN para desarrollo en local
	// strutura de DSN user:password@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local
	//cambia user por tu usuario de MySql y password por tu contrasenia de acceso a el puerto, por defecto uso 3306
	//handlers.ExecuteMigrations(DSN) // ejecut las migraciones crea la db si no existe y las tablas en esta si no existen
	//handlers.ImportDataFromCSVDB(DSN)

	// Inicializar Gin
	router := gin.Default()

	// Configurar las rutas
	routes.SetUpRoutes(router)

	// Iniciar el servidor
	router.Run(":8080")
}
