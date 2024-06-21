package activities_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOwnerActivitiesService() gin.HandlerFunc {
	return func(c *gin.Context) {

		user, err := Adapters.GetUserIdByToken(c)
		if err != nil {
			c.JSON(http.StatusNotFound, "Not user logged")
		}
		actividades, err2 := Adapters.GetActvitiesOwnerAdapter(user.ID)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar las actividades del usuario"})
			return
		}

		c.JSON(http.StatusOK, actividades)
	}
}
