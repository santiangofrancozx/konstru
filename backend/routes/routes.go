package routes

import (
	"awesomeKonstru/backend/middleware"
	"awesomeKonstru/backend/services"
	"awesomeKonstru/backend/services/sites"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	r := router
	//renders
	r.GET("/login", sites.RenderLoginTemplateService)
	//user
	r.POST("/userData", middleware.RequireAuth, services.GetUserInfoByTokenService())
	r.POST("/validateUser", services.ValidateLoginService())
	r.POST("/insertNewUser", services.InsertUserService())
	// consultas protegidas
	r.GET("/search", middleware.RequireAuth, sites.RenderBudgetTemplateService)
	r.GET("/consultaActividad", middleware.RequireAuth, services.GetActivityService())
	r.GET("/consultaApu", middleware.RequireAuth, services.GetActivityApuByActiityIdService())
	r.GET("/showInfo", middleware.RequireAuth, services.GetUserInfoByTokenService())
	r.GET("/logout", middleware.RequireAuth, services.LogoutUserByTokenService())
	r.POST("/insertInsumoUsuario", middleware.RequireAuth, services.InsertNewInsumoUserService())
}
