package projects_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Nombre   string
	Cantidad float64
	Precio   float64
	Total    float64
}

func GetAllProjectsActivitiesService() gin.HandlerFunc {
	return func(c *gin.Context) {
		idp := c.Query("id_proyecto")

		actividadesU, err := Adapters.SelectAllProjectsActivities(idp)

		var list []response

		for _, actividad := range actividadesU {
			actividad2, err := Adapters.SelectActivityByID(actividad.ID_actividad)
			if err != nil {
				return
			}
			totalM := actividad2.PrecioBase * actividad.Cantidad
			response := response{
				Nombre:   actividad2.Descripcion,
				Cantidad: actividad.Cantidad,
				Precio:   actividad2.PrecioBase,
				Total:    totalM,
			}
			list = append(list, response)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, list)
	}
}
