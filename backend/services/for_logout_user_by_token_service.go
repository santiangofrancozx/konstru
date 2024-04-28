package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutUserByTokenService() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Eliminar la cookie de autorización para cerrar sesión
		c.SetCookie("Authorization", "", -1, "", "", false, true)

		// Respond with a success message
		c.JSON(http.StatusOK, gin.H{
			"message": "Logout successful",
		})
	}
}
