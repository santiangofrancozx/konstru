package services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InsumosUsuarioRequest struct {
	IDInsumoUsuario string  `json:"id_insumo_usuario"`
	Descripcion     string  `json:"descripcion"`
	Unidad          string  `json:"unidad"`
	PrecioBase      float64 `json:"precio_base"`
}
type InsumosUsuarioResponse struct {
	ServiceUsed       string  `json:"service-used"`
	IDInsumoUsuario   string  `json:"id_insumo_usuario"`
	Descripcion       string  `json:"descripcion"`
	Unidad            string  `json:"unidad"`
	PrecioBase        float64 `json:"precio_base"`
	UsuarioIdentifier string  `json:"usuario_identifier"`
}

func InsertNewInsumoUserService() gin.HandlerFunc {
	return func(c *gin.Context) {
		var InsumosUsuarioRequest InsumosUsuarioRequest
		if err := c.BindJSON(&InsumosUsuarioRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user := Adapters.GetUserIdByToken(c)

		insumoUpload := models.Insumos_Usuario{
			ID_Insumo_Usuario: InsumosUsuarioRequest.IDInsumoUsuario,
			Usuario_ID:        user.ID,
			Descripcion:       InsumosUsuarioRequest.Descripcion,
			Unidad:            InsumosUsuarioRequest.Unidad,
			PrecioBase:        InsumosUsuarioRequest.PrecioBase,
		}

		errI := Adapters.InserteIntoInsumoUsuarioAdapter(insumoUpload)
		if errI != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errI.Error()})
			return // Aquí se detiene la ejecución de la función si hay un error
		}

		// Si no hay error, continuamos con la respuesta exitosa
		insumoResponse := InsumosUsuarioResponse{
			ServiceUsed:       "InsertNewInsumoUserService",
			IDInsumoUsuario:   InsumosUsuarioRequest.IDInsumoUsuario,
			Descripcion:       InsumosUsuarioRequest.Descripcion,
			Unidad:            InsumosUsuarioRequest.Unidad,
			PrecioBase:        InsumosUsuarioRequest.PrecioBase,
			UsuarioIdentifier: user.Email,
		}
		c.JSON(http.StatusOK, insumoResponse)
	}
}
