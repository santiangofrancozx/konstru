package services

import (
	"awesomeKonstru/backend/handlers/Query"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInsumoService(DSN string) gin.HandlerFunc {
	return func(c *gin.Context) {

		ID := c.Query("id")

		insumo, _ := Query.SelectInsumoByID(DSN, ID)

		insumos, err := Query.SelectInsumoByNombre(DSN, ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar las actividades por nombre"})
			return
		}

		insumos = append(insumos, insumo)
		response := jsonFormat{
			ServiceUsed: "GetInsumoService",
			Data:        insumos,
		}

		c.JSON(http.StatusOK, response)
	}

}
