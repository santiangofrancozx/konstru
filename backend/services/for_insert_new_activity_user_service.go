package services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ActivityUserRequest struct {
	IDActivityUser string  `json:"IDActivityUser"`
	Descripcion    string  `json:"descripcion"`
	Unidad         string  `json:"unidad"`
	PrecioBase     float64 `json:"precio_base"`
}
type ActivityUserResponse struct {
	ServiceUsed string `json:"service-used"`
	ActivityUserRequest
	UsuarioIdentifier string `json:"usuario_identifier"`
}

func InsertNewActivityUserService() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ActivityUserRequest ActivityUserRequest
		if err := c.BindJSON(&ActivityUserRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user := Adapters.GetUserIdByToken(c)
		activityUpload := models.Actividad_Usuario{
			ID_Actividad_Usuario: ActivityUserRequest.IDActivityUser,
			Usuario_ID:           user.ID,
			Descripcion:          ActivityUserRequest.Descripcion,
			Unidad:               ActivityUserRequest.Unidad,
			PrecioBase:           ActivityUserRequest.PrecioBase,
		}
		errI := Adapters.InsertIntoActivityUserAdapter(activityUpload)
		if errI != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errI.Error()})
			return // Aquí se detiene la ejecución de la función si hay un error
		}

		activityResponse := ActivityUserResponse{
			ServiceUsed:         "InsertNewActivityUserService",
			ActivityUserRequest: ActivityUserRequest,
			UsuarioIdentifier:   user.Email,
		}
		c.JSON(http.StatusOK, activityResponse)
	}

}
