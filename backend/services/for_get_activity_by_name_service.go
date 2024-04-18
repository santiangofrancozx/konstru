package services

import (
	"awesomeKonstru/backend/handlers/Query"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetActivityByNameService(DSN string) gin.HandlerFunc {
	//DSN := config.DB_DSN
	return func(context *gin.Context) {
		c := context

		ID := c.Query("id")

		actividades, err := Query.SelectActividadByNombre(DSN, ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := jsonFormat{
			ServiceUsed: "GetActivityByNameService",
			Data:        actividades,
		}
		c.JSON(http.StatusOK, response)
	}
}
