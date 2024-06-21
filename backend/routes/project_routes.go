package routes

import (
	"awesomeKonstru/backend/middleware"
	"awesomeKonstru/backend/services/projects_services"
	"github.com/gin-gonic/gin"
)

func SetUpProjectRoutes(router *gin.Engine) {
	r := router
	r.POST("/projects/create", middleware.RequireAuth, projects_services.InsertNewProjectService())
	r.POST("/projects/activities/create", middleware.RequireAuth, projects_services.InsertNewProjectActivitiesService())
	r.GET("/projects/get", projects_services.GetUserProjectsByTokenService())
	r.GET("/projects/activities/get", projects_services.GetAllProjectsActivitiesService())
	r.DELETE("/projects/activities/delete", projects_services.DeleteProhectActivitiesService())
	r.DELETE("/projects/delete", projects_services.DeleteProjectsByProjectIdOnCascade())
	r.PUT("/projects/activities/update", projects_services.UpdateProjectActivitiesService())
}
