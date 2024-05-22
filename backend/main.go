package main

import (
	"awesomeKonstru/backend/config"
	"net/http"

	//"awesomeKonstru/backend/handlers/Connection-Migrates/Migrates"
	"awesomeKonstru/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// strutura de DSN user:password@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local
	config.SetDSN("root:safraroot@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local")
	//cambia user por tu usuario de MySql y password por tu contrasenia de acceso a el puerto, por defecto uso 3306
	//Migrates.MakeMigrations(Migrates.ExecuteMigrations()) // ejecuta las migraciones crea la db si no existe y las tablas en esta si no existe
	//Migrates.ImportDataFromCSVDB()

	router := gin.Default()
	gin.SetMode(gin.DebugMode) // o gin.ReleaseMode
	router.Static("/static", "../frontend/static")

	router.Use(corsMiddleware())

	// Configurar las rutas
	routes.SetUpRoutes(router)

	// Iniciar el servidor
	router.Run(":8082")

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // Asegúrate de permitir credenciales
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}
