package main

import (
	"awesomeKonstru/backend/config"
	//handlers "awesomeKonstru/backend/handlers/Connection-Migrates/Migrates"
	"awesomeKonstru/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//DSN := "" // Define aqui tu DSN para desarrollo en local
	// strutura de DSN user:password@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local
	config.SetDSN("root:safraroot@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local")
	//cambia user por tu usuario de MySql y password por tu contrasenia de acceso a el puerto, por defecto uso 3306
	//handlers.ExecuteMigrations(config.DB_DSN) // ejecut las migraciones crea la db si no existe y las tablas en esta si no existen
	//handlers.ImportDataFromCSVDB(config.DB_DSN)

	// Inicializar Gin

	router := gin.Default()
	gin.SetMode(gin.DebugMode) // o gin.ReleaseMode
	router.Static("/static", "../frontend/static")

	// Configurar las rutas
	routes.SetUpRoutes(router)

	// Iniciar el servidor
	router.Run(":8082")
}
