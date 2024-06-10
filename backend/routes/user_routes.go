package routes

import (
	"awesomeKonstru/backend/middleware"
	"awesomeKonstru/backend/services/user_services"
	"github.com/gin-gonic/gin"
)

func SetUpUserRoutes(router *gin.Engine) {
	r := router
	r.GET("/user/data", middleware.RequireAuth, user_services.GetUserInfoByTokenService())
	r.POST("/user/validate", user_services.ValidateLoginService())
	r.POST("/user/create", user_services.InsertUserService())
	r.GET("/user/logout", middleware.RequireAuth, user_services.LogoutUserByTokenService())
}
