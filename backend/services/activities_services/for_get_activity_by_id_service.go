package activities_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActivityByIdService() gin.HandlerFunc {
	return func(c *gin.Context) {

		ID := c.Query("id")

		actividad, err := Adapters.SelectActivityByID(ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := JsonFormat{
			ServiceUsed: "GetActivityByIdSerive",
			Data:        actividad,
		}
		c.JSON(http.StatusOK, response)
	}
}
