package insumos_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InsumosRequest struct {
	ID          string  `json:"ID"`
	Descripcion string  `json:"descripcion"`
	Unidad      string  `json:"unidad"`
	PrecioBase  float64 `json:"precio_base"`
}
type InsumosResponse struct {
	ServiceUsed string `json:"service-used"`
	InsumosRequest
	//UsuarioIdentifier string `json:"usuario_identifier"`
}

func InsertNewInsumoService() gin.HandlerFunc {
	return func(c *gin.Context) {
		var InsumosRequest InsumosRequest
		if err := c.BindJSON(&InsumosRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, _ := Adapters.GetUserIdByToken(c)
		ID := user.ID

		insumoUpload := models.Insumo{
			ID:          InsumosRequest.ID,
			Descripcion: InsumosRequest.Descripcion,
			Unidad:      InsumosRequest.Unidad,
			PrecioBase:  InsumosRequest.PrecioBase,
			CreatedBy:   &ID,
		}

		errI := Adapters.InsertIntoInsumoAdapter(insumoUpload)
		if errI != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errI.Error()})
			return // Aquí se detiene la ejecución de la función si hay un error
		}

		// Si no hay error, continuamos con la respuesta exitosa
		insumoResponse := InsumosResponse{
			ServiceUsed:    "InsertNewInsumoService",
			InsumosRequest: InsumosRequest,
			//UsuarioIdentifier: user.Email,
		}
		c.JSON(http.StatusOK, insumoResponse)
	}
}
