package services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var InsumosUsuarioRequest struct {
	IDInsumoUsuario string  `json:"id_insumo_usuario"`
	UsuarioID       int     `json:"usuario_id"`
	Descripcion     string  `json:"descripcion"`
	Unidad          string  `json:"unidad"`
	PrecioBase      float64 `json:"precio_base"`
}

func InsertNewInsumoUserService() gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := c.BindJSON(&InsumosUsuarioRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		insumoUpload := models.Insumos_Usuario{
			ID_Insumo_Usuario: InsumosUsuarioRequest.IDInsumoUsuario,
			Usuario_ID:        InsumosUsuarioRequest.UsuarioID,
			Descripcion:       InsumosUsuarioRequest.Descripcion,
			Unidad:            InsumosUsuarioRequest.Unidad,
			PrecioBase:        InsumosUsuarioRequest.PrecioBase,
		}

		// Buscar el usuario por su correo electrónico
		_, err := Adapters.InserteIntoInsumoUsuarioAdapter(insumoUpload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el usuario"})
			return
		}
		c.JSON(http.StatusOK, InsumosUsuarioRequest)
	}
}
