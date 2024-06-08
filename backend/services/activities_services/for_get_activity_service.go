package activities_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonFormat struct {
	ServiceUsed string      `json:"service_used"`
	Data        interface{} `json:"data"`
}

func GetActivityService() gin.HandlerFunc {
	return func(c *gin.Context) {

		ID := c.Query("id")

		actividad, _ := Adapters.SelectActivityByID(ID)

		actividades, err := Adapters.SelectActividadByNombre(ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar las actividades por nombre"})
			return
		}

		actividades = append(actividades, actividad)

		c.JSON(http.StatusOK, actividades)
	}
}
