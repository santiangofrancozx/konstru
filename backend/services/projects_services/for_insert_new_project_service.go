package projects_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProyectoRequest struct {
	IDProyecto  string  `json:"id_proyecto"`
	Name        string  `json:"name"`
	Descripcion string  `json:"descripcion"`
	Valor       float64 `json:"valor"`
	TipoObra    string  `json:"tipo_obra"`
}
type ProjectResponse struct {
	ServiceUsed string `json:"service-used"`
	ProyectoRequest
	UserId string `json:"user-id"`
}

func InsertNewProjectService() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ProyectoRequest ProyectoRequest
		if err := c.BindJSON(&ProyectoRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := Adapters.GetUserIdByToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}

		projectUpload := models.Proyectos{
			IDProyecto:  ProyectoRequest.IDProyecto,
			Name:        ProyectoRequest.Name,
			Descripcion: ProyectoRequest.Descripcion,
			Valor:       ProyectoRequest.Valor,
			TipoObra:    ProyectoRequest.TipoObra,
			UsuarioID:   user.ID,
		}
		errI := Adapters.InsertNewProjectAdapter(projectUpload)
		if errI != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errI.Error()})
			return // Aquí se detiene la ejecución de la función si hay un error
		}

		// Si no hay error, continuamos con la respuesta exitosa
		ProjectResponse := ProjectResponse{
			ServiceUsed:     "InsertNewProjectService",
			ProyectoRequest: ProyectoRequest,
			UserId:          user.ID,
		}
		c.JSON(http.StatusOK, ProjectResponse)

	}

}
