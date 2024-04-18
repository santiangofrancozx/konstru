package services

import (
	"awesomeKonstru/backend/handlers/Query"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func ValidateLoginService() gin.HandlerFunc {
	return func(context *gin.Context) {
		c := context
		var requestData struct {
			Email string `json:"email"`
			Pass  string `json:"pass"`
		}
		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Buscar el usuario por su correo electrónico
		usuario, err := Query.SelectUserByUsername(requestData.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el usuario"})
			return
		}

		// Verificar si el usuario fue encontrado y si la contraseña coincide
		if verifyPassword(usuario.Password, requestData.Pass) {

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub":   usuario.ID,
				"email": usuario.Email,
				"name":  usuario.Nombre,
				"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Failed to create token",
				})
				return
			}

			// Respond
			c.SetSameSite(http.SameSiteLaxMode)
			c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

			c.JSON(http.StatusOK, tokenString)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid email or password",
			})
		}
	}
}

func verifyPassword(storedPassword, providedPassword string) bool {
	// Comparar la contraseña encriptada proporcionada con la contraseña encriptada almacenada
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}
