package Adapters

import (
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserIdByToken(c *gin.Context) (models.Usuario, error) {
	token, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		return models.Usuario{}, err
	}
	tokenString, err1 := DecodeJWTToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		return models.Usuario{}, err1
	}
	email, ok := tokenString["email"].(string)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return models.Usuario{}, err
	}
	user, err := SelectUserByUsername(email)
	return user, nil
}
