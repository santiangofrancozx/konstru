package auth_services

import (
	"awesomeKonstru/backend/middleware"
	"github.com/gin-gonic/gin"
)

func CheckTokenService() gin.HandlerFunc {
	return func(context *gin.Context) {
		middleware.RequireAuth(context)
	}
}
