package activities_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteActivityByIdService() gin.HandlerFunc {
	return func(c *gin.Context) {
		activityId := c.Query("id")

		err := Adapters.ActivityByID(activityId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la actividad"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Activity deleted succesfully"})
	}
}
