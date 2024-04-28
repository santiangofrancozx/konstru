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
	r.GET("/register", sites.RenderRegisterTemplateService)
	r.POST("/userData", middleware.RequireAuth, services.GetUserByEmailService())
	r.POST("/validateUser", services.ValidateLoginService())
	r.POST("/registerPost", services.InsertUserService())
	// Utiliza el middleware para proteger la ruta '/search'
	r.GET("/search", middleware.RequireAuth, sites.RenderBudgetTemplateService)
	r.GET("/consultaActividad", middleware.RequireAuth, services.GetActivityService())
	r.GET("/consultaApu", middleware.RequireAuth, services.GetActivityApuByActiityIdService())
	r.GET("/showInfo", middleware.RequireAuth, services.GetUserInfoByTokenService())
}
