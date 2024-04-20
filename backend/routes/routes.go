package routes

import (
	"awesomeKonstru/backend/middleware"
	"awesomeKonstru/backend/services"
	"awesomeKonstru/backend/services/sites"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	r := router
	r.GET("/login", sites.RenderLoginTemplateService)
	r.POST("/userData", services.GetUserByEmailService())
	r.POST("/validateUser", services.ValidateLoginService())
	// Utiliza el middleware para proteger la ruta '/search'
	r.GET("/search", middleware.RequireAuth, sites.RenderBudgetTemplateService)
	r.GET("/consultaActividad", middleware.RequireAuth, services.GetActivityService())
	r.GET("/consultaApu", middleware.RequireAuth, services.GetActivityApuByActiityIdService())

}
