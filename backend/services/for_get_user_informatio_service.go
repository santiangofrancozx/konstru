package services

import (
	"awesomeKonstru/backend/handlers/Query"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserByEmailService() gin.HandlerFunc {
	return func(context *gin.Context) {
		c := context
		var requestData struct {
			Email string `json:"email"`
		}
		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		usuario, err := Query.SelectUserByUsername(requestData.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el usuario" + requestData.Email})
			return
		}
		c.JSON(http.StatusOK, usuario)

	}
}
