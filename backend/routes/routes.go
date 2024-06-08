package routes

import (
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	// Render routes
	//SetUpRenderRoutes(router)
	// Auth routes
	SetUpAuthRoutes(router)
	// User routes
	SetUpUserRoutes(router)
	// Insumos routes
	SetUpInsumosRoutes(router)
	// Activity routes
	SetUpActivityRoutes(router)
	// Project routes
	SetUpProjectRoutes(router)
	// Apu routes
	SetUpApuRoutes(router)
}
