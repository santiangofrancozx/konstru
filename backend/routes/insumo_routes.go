package routes

import (
	"awesomeKonstru/backend/middleware"
	"awesomeKonstru/backend/services/insumos_services"
	"github.com/gin-gonic/gin"
)

func SetUpInsumosRoutes(router *gin.Engine) {
	r := router
	r.POST("/insumos", middleware.RequireAuth, insumos_services.InsertNewInsumoService())
	r.GET("/insumos", insumos_services.GetInsumoService())
}
