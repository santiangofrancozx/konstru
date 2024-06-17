package projects_services

import (
	"awesomeKonstru/backend/handlers/Adapters"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProjectsActivitiesService() gin.HandlerFunc {
	return func(c *gin.Context) {
		actividadesU, err := Adapters.SelectAllProjectsActivities()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, actividadesU)
	}
}
	