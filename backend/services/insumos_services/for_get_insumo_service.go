package insumos_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInsumoService() gin.HandlerFunc {
	return func(c *gin.Context) {

		ID := c.Query("id")

		insumo, _ := Adapters.SelectInsumoByID(ID)

		insumos, err := Adapters.SelectInsumoByNombre(ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar las actividades por nombre"})
			return
		}

		insumos = append(insumos, insumo)
		//response := get_activity_services.JsonFormat{
		//	ServiceUsed: "GetInsumoService",
		//	Data:        insumos,
		//}

		c.JSON(http.StatusOK, insumos)
	}

}
