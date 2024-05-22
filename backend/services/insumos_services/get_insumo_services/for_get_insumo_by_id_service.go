package get_insumo_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/services/activities_services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInsumoByIDService() gin.HandlerFunc {
	return func(context *gin.Context) {
		c := context
		id := c.Query("id")
		insumo, err := Adapters.SelectInsumoByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := activities_services.jsonFormat{
			ServiceUsed: "GetInsumoByIdService",
			Data:        insumo,
		}
		c.JSON(http.StatusOK, response)
	}

}