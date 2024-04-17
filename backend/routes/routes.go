package routes

import (
	"awesomeKonstru/backend/config"
	"awesomeKonstru/backend/services"
	"awesomeKonstru/backend/services/sites"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	r := router
	//r.Static("./frontend/static", "./frontend/static")
	//GET
	//r.GET("/", handler)
	r.GET("/login", sites.RenderLoginTemplateService)
	r.GET("/search", sites.RenderBudgetTemplateService)
	r.GET("/consultaActividad", services.GetActivityService(config.DB_DSN))
	//r.GET("/register", RegisterHandler)

	//POST
	//r.POST("/registerPost", registerPostHandler)
	//r.POST("/validarLogin", loginPostHandler)
}
