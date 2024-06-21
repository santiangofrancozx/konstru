package projects_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ActivitiesUpdate struct {
	Id_activity string  `json:"Id_activity"`
	Quantity    float64 `json:"Quantity"`
}

type UpdateProjectActivitiesRequest struct {
	Id_project string             `json:"Id_project"`
	Activities []ActivitiesUpdate `json:"Activities"`
}

func UpdateProjectActivitiesService() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req UpdateProjectActivitiesRequest
		var modelsU []models.Proyectos_actividades
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
			return
		}
		project, err := Adapters.GetProjectByIdAdapter(req.Id_project)
		if err != nil {
			c.JSON(http.StatusNotFound, "Project not found")
			return
		}

		user, errU := Adapters.GetUserIdByToken(c)
		if errU != nil {
			c.JSON(http.StatusUnauthorized, "User not logged")
			return
		}
		if project.UsuarioID != user.ID {
			c.JSON(http.StatusUnauthorized, "The project does not belong you")
			return
		}

		if len(req.Activities) > 0 {
			for _, activity := range req.Activities {
				upload := models.Proyectos_actividades{
					ID_proyecto:  req.Id_project,
					ID_actividad: activity.Id_activity,
					Cantidad:     activity.Quantity,
				}
				modelsU = append(modelsU, upload)
			}
		}

		succes := Adapters.UpdateProjectActivitiesAdapter(modelsU, req.Id_project)
		if succes != nil {
			c.JSON(http.StatusInternalServerError, succes)
		}

		c.JSON(http.StatusOK, modelsU)
	}
}
