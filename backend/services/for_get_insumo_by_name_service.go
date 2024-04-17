package services

import (
	"awesomeKonstru/backend/handlers/Query"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInsumoByNameService(DSN string) gin.HandlerFunc {
	return func(context *gin.Context) {
		c := context
		id := c.Query("id")
		insumos, err := Query.SelectInsumoByNombre(DSN, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := jsonFormat{
			ServiceUsed: "GetInsumoByNameService",
			Data:        insumos,
		}
		c.JSON(http.StatusOK, response)
	}

}
