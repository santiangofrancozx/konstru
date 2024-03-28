package routes

import (
	"awesomeKonstru/backend/handlers/routes"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) {
	r := router
	r.Static("./frontend/static", "./frontend/static")
	//GET
	//r.GET("/", handler)
	r.GET("/login", routes.LoginHandler)
	//r.GET("/register", RegisterHandler)

	//POST
	//r.POST("/registerPost", registerPostHandler)
	//r.POST("/validarLogin", loginPostHandler)
}
