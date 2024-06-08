package routes

import (
	"awesomeKonstru/backend/services/auth_services"
	"github.com/gin-gonic/gin"
)

func SetUpAuthRoutes(router *gin.Engine) {
	r := router
	r.GET("/auth/check-token", auth_services.CheckTokenService())
}
