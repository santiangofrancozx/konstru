package get_insumo_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInsumoByNameService() gin.HandlerFunc {
	return func(context *gin.Context) {
		c := context
		id := c.Query("id")
		insumos, err := Adapters.SelectInsumoByNombre(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//response := activities_services.jsonFormat{
		//	ServiceUsed: "GetInsumoByNameService",
		//	Data:        insumos,
		//}
		c.JSON(http.StatusOK, insumos)
	}

}
