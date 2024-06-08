package routes

import (
	"awesomeKonstru/backend/middleware"
	"awesomeKonstru/backend/services/activities_services"
	"github.com/gin-gonic/gin"
)

func SetUpActivityRoutes(router *gin.Engine) {
	r := router
	r.POST("/activities/create", middleware.RequireAuth, activities_services.InsertNewActivityService())
	r.POST("/activities/user/create", middleware.RequireAuth, activities_services.InsertNewActivityUserService())
	r.GET("/activities/get", activities_services.GetActivityService())
	r.GET("/activities/user/get", activities_services.GetActivitiesService())
	r.GET("/activities/apu/get", activities_services.GetActivityApuByActiityIdService())
}
