package services

import (
	"awesomeKonstru/backend/handlers/Query"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetActivityByIdService() gin.HandlerFunc {
	return func(context *gin.Context) {
		c := gin.Context{}

		ID := c.Query("id")

		actividad, err := Query.SelectActivityByID(ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := jsonFormat{
			ServiceUsed: "GetActivityByIdSerive",
			Data:        actividad,
		}
		c.JSON(http.StatusOK, response)
	}
}
