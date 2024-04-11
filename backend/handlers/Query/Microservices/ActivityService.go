package Microservices

import (
	"awesomeKonstru/backend/handlers/Query"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleActivityInsumo(DSN string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var requestData struct {
			ID string `json:"id"`
		}
		// Bind JSON a la estructura requestData
		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		actividad, _ := Query.SelectActivityByID(DSN, requestData.ID)

		actividades, err := Query.SelectActividadByNombre(DSN, requestData.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar las actividades por nombre"})
			return
		}

		actividades = append(actividades, actividad)

		c.JSON(http.StatusOK, actividades)
	}
}

func HandleActividadByID(c *gin.Context, DSN, id string) {
	actividad, err := Query.SelectActivityByID(DSN, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actividad)
}

func HandleActividadByNombre(c *gin.Context, DSN, id string) {
	actividades, err := Query.SelectActividadByNombre(DSN, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actividades)
}
