package activities_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ActivityUserRequest struct {
	Descripcion string  `json:"descripcion"`
	Unidad      string  `json:"unidad"`
	PrecioBase  float64 `json:"precio_base"`
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
		user, err := Adapters.GetUserIdByToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		activityUpload := models.Actividad_Usuario{
			Usuario_ID:  user.ID,
			Descripcion: ActivityUserRequest.Descripcion,
			Unidad:      ActivityUserRequest.Unidad,
			PrecioBase:  ActivityUserRequest.PrecioBase,
		}
		errI, _ := Adapters.InsertIntoActivityUserAdapter(activityUpload)
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
