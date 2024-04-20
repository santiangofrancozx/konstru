package services

import (
	"awesomeKonstru/backend/handlers/Query"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInsumoByIDService() gin.HandlerFunc {
	return func(context *gin.Context) {
		c := context
		id := c.Query("id")
		insumo, err := Query.SelectInsumoByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := jsonFormat{
			ServiceUsed: "GetInsumoByIdService",
			Data:        insumo,
		}
		c.JSON(http.StatusOK, response)
	}

}
