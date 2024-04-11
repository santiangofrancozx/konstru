package routes

import (
	"awesomeKonstru/backend/config"
	"awesomeKonstru/backend/handlers/Query/Microservices"
	"awesomeKonstru/backend/handlers/sites"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	r := router
	//r.Static("./frontend/static", "./frontend/static")
	//GET
	//r.GET("/", handler)
	r.GET("/login", sites.LoginHandler)
	r.GET("/search", sites.SearchRenderHandler)
	r.POST("/consultaActividad", Microservices.HandleActivityInsumo(config.DB_DSN))
	//r.GET("/register", RegisterHandler)

	//POST
	//r.POST("/registerPost", registerPostHandler)
	//r.POST("/validarLogin", loginPostHandler)
}
