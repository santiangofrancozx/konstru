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
	r.POST("/userData", middleware.RequireAuth, services.GetUserInfoByTokenService())
	r.POST("/validateUser", services.ValidateLoginService())
	r.POST("/insertNewUser", services.InsertUserService())
	// Utiliza el middleware para proteger la ruta '/search'
	r.GET("/search", sites.RenderBudgetTemplateService)
	r.GET("/consultaActividad", middleware.RequireAuth, services.GetActivityService())
	r.GET("/consultaApu", services.GetActivityApuByActiityIdService())
	r.GET("/showInfo", middleware.RequireAuth, services.GetUserInfoByTokenService())
	r.GET("/logout", middleware.RequireAuth, services.LogoutUserByTokenService())
	r.POST("insertInsumoUsuario", services.InsertNewInsumoUserService())
}
