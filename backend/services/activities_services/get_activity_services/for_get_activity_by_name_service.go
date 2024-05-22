package get_activity_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetActivityByNameService() gin.HandlerFunc {
	//DSN := config.DB_DSN
	return func(context *gin.Context) {
		c := context

		ID := c.Query("id")

		actividades, err := Adapters.SelectActividadByNombre(ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := JsonFormat{
			ServiceUsed: "GetActivityByNameService",
			Data:        actividades,
		}
		c.JSON(http.StatusOK, response)
	}
}
