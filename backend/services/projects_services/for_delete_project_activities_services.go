package projects_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProyectoActivityDeleteRequest struct {
	IDProyecto   string   `json:"id_proyecto"`
	IDActivities []string `json:"id_activities"`
}

func DeleteProhectActivitiesService() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ProyectoActivityDeleteRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
			return
		}

		// Validar que se envíen IDs de actividades
		if len(req.IDActivities) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No activity IDs provided"})
			return
		}

		// Obtener el proyecto por ID
		project, err := Adapters.GetProjectByIdAdapter(req.IDProyecto)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}

		// Obtener el usuario autenticado
		user, errU := Adapters.GetUserIdByToken(c)
		if errU != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged"})
			return
		}

		// Verificar si el proyecto pertenece al usuario autenticado
		if project.UsuarioID != user.ID {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Project does not belong to this user"})
			return
		}

		// Llamar al adaptador para eliminar las actividades del proyecto
		err = Adapters.DeleteProjectActivitiesAdapter(req.IDProyecto, req.IDActivities)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project activities"})
			return
		}

		// Respuesta exitosa
		c.JSON(http.StatusOK, gin.H{"message": "Project activities deleted successfully"})
	}
}
