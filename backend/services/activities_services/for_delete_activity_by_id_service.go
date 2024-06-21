package activities_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteActivityByIdService() gin.HandlerFunc {
	return func(c *gin.Context) {
		activityId := c.Query("id")

		user, err := Adapters.GetUserIdByToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "No tienes autorizacion")
			return
		}

		err2 := Adapters.ActivityByID(activityId, user.ID)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la actividad"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Activity deleted succesfully"})
	}
}
