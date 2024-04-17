package services

import (
	"awesomeKonstru/backend/handlers/Query"
	"github.com/gin-gonic/gin"
	"net/http"
)

type jsonFormat struct {
	ServiceUsed string      `json:"service_used"`
	Data        interface{} `json:"data"`
}

func GetActivityService(DSN string) gin.HandlerFunc {
	return func(c *gin.Context) {

		ID := c.Query("id")

		actividad, _ := Query.SelectActivityByID(DSN, ID)

		actividades, err := Query.SelectActividadByNombre(DSN, ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar las actividades por nombre"})
			return
		}

		actividades = append(actividades, actividad)
		response := jsonFormat{
			ServiceUsed: "GetActivityService",
			Data:        actividades,
		}

		c.JSON(http.StatusOK, response)
	}
}
