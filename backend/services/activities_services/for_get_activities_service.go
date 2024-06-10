package activities_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetActivitiesService() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := Adapters.GetUserIdByToken(c)

		actividadesU, err2 := Adapters.SelectActivityByUserID(user.ID)

		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar las actividades"})
			return
		}

		c.JSON(http.StatusOK, actividadesU)
	}
}
