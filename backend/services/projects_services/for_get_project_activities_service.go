package projects_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Descripcion string
	Cantidad    float64
	PrecioBase  float64
	Total       float64
	ID          string
}

func GetAllProjectsActivitiesService() gin.HandlerFunc {
	return func(c *gin.Context) {
		idp := c.Query("id_proyecto")
		project, err := Adapters.GetProjectByIdAdapter(idp)
		if err != nil {
			c.JSON(http.StatusNotFound, "Project not found")
		}

		user, errU := Adapters.GetUserIdByToken(c)
		if errU != nil {
			c.JSON(http.StatusUnauthorized, "User not logged")
		}
		if project.UsuarioID != user.ID {
			c.JSON(http.StatusUnauthorized, "The project does not belong you")
		}

		actividadesU, err := Adapters.SelectAllProjectsActivities(idp)

		var list []response

		for _, actividad := range actividadesU {
			actividad2, err := Adapters.SelectActivityByID(actividad.ID_actividad)
			if err != nil {
				return
			}
			totalM := actividad2.PrecioBase * actividad.Cantidad
			response := response{
				Descripcion: actividad2.Descripcion,
				Cantidad:    actividad.Cantidad,
				PrecioBase:  actividad2.PrecioBase,
				Total:       totalM,
				ID:          actividad2.ID,
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
