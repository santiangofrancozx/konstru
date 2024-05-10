package services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ActivityRequest struct {
	ID          string  `json:"ID"`
	Descripcion string  `json:"descripcion"`
	Unidad      string  `json:"unidad"`
	PrecioBase  float64 `json:"precio_base"`
}
type ActivityResponse struct {
	ServiceUsed string `json:"service-used"`
	ActivityRequest
	//UsuarioIdentifier string `json:"usuario_identifier"`
}

func InsertNewActivityService() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ActivityRequest ActivityRequest
		if err := c.BindJSON(&ActivityRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//user := Adapters.GetUserIdByToken(c)
		activityUpload := models.Actividad{
			ID:          ActivityRequest.ID,
			Descripcion: ActivityRequest.Descripcion,
			Unidad:      ActivityRequest.Unidad,
			PrecioBase:  ActivityRequest.PrecioBase,
		}
		errI := Adapters.InsertIntoActivityAdapter(activityUpload)
		if errI != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errI.Error()})
			return // Aquí se detiene la ejecución de la función si hay un error
		}

		activityResponse := ActivityResponse{
			ServiceUsed:     "InsertNewActivityService",
			ActivityRequest: ActivityRequest,
			//UsuarioIdentifier: user.Email,
		}
		c.JSON(http.StatusOK, activityResponse)
	}

}
