package projects_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Activities struct {
	ID_actividad string  `json:"ID_Actividad"`
	Cantidad     float64 `json:"Cantidad"`
}

type ProyectoActivityRequest struct {
	IDProyecto string       `json:"id_proyecto"`
	Activities []Activities `json:"activities"`
}

//dump Json postman test
//{
//  "id_proyecto": "12345",
//  "activities": [
//    {
//      "ID_actividad": "A1",
//      "cantidad": 100.5
//    },
//    {
//      "ID_actividad": "A2",
//      "cantidad": 200.75
//    }
//  ]
//}

func InsertNewProjectActivitiesService() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ProyectoActivityRequest ProyectoActivityRequest
		if err := c.BindJSON(&ProyectoActivityRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Security fisrt steps, verified project owner
		project, err := Adapters.GetProjectByIdAdapter(ProyectoActivityRequest.IDProyecto)
		if err != nil {
			c.JSON(http.StatusNotFound, "Project not found")
		}

		user, errU := Adapters.GetUserIdByToken(c)
		if errU != nil {
			c.JSON(http.StatusUnauthorized, "User not logged")
		}
		var projectUpload []models.Proyectos_actividades
		if project.UsuarioID == user.ID {
			for _, activity := range ProyectoActivityRequest.Activities {
				projectUpload = append(projectUpload, models.Proyectos_actividades{
					ID_proyecto:  ProyectoActivityRequest.IDProyecto,
					ID_actividad: activity.ID_actividad,
					Cantidad:     activity.Cantidad,
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, "project not found or does not belong to this user")
			return
		}

		errI := Adapters.InsertNewProjectActivitiesAdapter(projectUpload)
		if errI != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errI.Error()})
			return // Aquí se detiene la ejecución de la función si hay un error
		}

		// Si no hay error, continuamos con la respuesta exitosa
		c.JSON(http.StatusOK, projectUpload)

	}

}
