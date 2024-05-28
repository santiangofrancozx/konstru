package routes

import (
	"awesomeKonstru/backend/middleware"
	"awesomeKonstru/backend/services/activities_services/create_activity_services"
	"awesomeKonstru/backend/services/activities_services/get_activity_services"
	"awesomeKonstru/backend/services/apu_services"
	"awesomeKonstru/backend/services/auth_services"
	"awesomeKonstru/backend/services/insumos_services/create_insumo_services"
	"awesomeKonstru/backend/services/insumos_services/get_insumo_services"
	"awesomeKonstru/backend/services/projects_services/create_project_services"
	"awesomeKonstru/backend/services/projects_services/get_project_services"
	"awesomeKonstru/backend/services/sites"
	"awesomeKonstru/backend/services/user_services/get_user_services"
	"awesomeKonstru/backend/services/user_services/insert_user_services"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	r := router
	//renders:
	r.GET("/login", sites.RenderLoginTemplateService)
	r.GET("/home", middleware.RequireAuth, sites.RenderHomePageService)
	//user:
	r.GET("/checkToken", auth_services.CheckTokenService())
	r.GET("/userData", middleware.RequireAuth, get_user_services.GetUserInfoByTokenService())
	r.POST("/validateUser", get_user_services.ValidateLoginService())
	r.POST("/insertNewUser", insert_user_services.InsertUserService())
	r.POST("/insertInsumo", middleware.RequireAuth, create_insumo_services.InsertNewInsumoService())
	r.POST("/insertActivity", middleware.RequireAuth, create_activity_services.InsertNewActivityService())
	r.POST("/insertActividadUsuario", middleware.RequireAuth, create_activity_services.InsertNewActivityUserService())
	r.POST("/insertProject", middleware.RequireAuth, create_project_services.InsertNewProjectService())
	// consultas protegidas:
	r.GET("/projects", get_project_services.GetUserProjectsByTokenService())
	r.GET("/search", middleware.RequireAuth, sites.RenderBudgetTemplateService)
	r.GET("/consultaActividad", get_activity_services.GetActivityService())
	r.GET("/consultaApu", get_activity_services.GetActivityApuByActiityIdService())
	r.GET("/consultaInsumo", get_insumo_services.GetInsumoService())
	r.GET("/showInfo", middleware.RequireAuth, get_user_services.GetUserInfoByTokenService())
	r.GET("/logout", middleware.RequireAuth, get_user_services.LogoutUserByTokenService())
	r.POST("/createapu", apu_services.InsertNewApuAndActivityUserService())
}
