package routes

import (
	"awesomeKonstru/backend/services/apu_services"
	"github.com/gin-gonic/gin"
)

func SetUpApuRoutes(router *gin.Engine) {
	r := router
	r.POST("/apu/create", apu_services.InsertNewApuAndActivityUserService())
}
