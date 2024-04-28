package services

import (
	"awesomeKonstru/backend/handlers/Query"
	"awesomeKonstru/backend/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func InsertUserService() gin.HandlerFunc {
	return func(c *gin.Context) {

		nombres := c.PostForm("nombres")
		apellidos := c.PostForm("apellidos")
		email := c.PostForm("email")
		password := c.PostForm("pass")

		// Codificacion de la password
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Error al codificar la contraseña:", err)
		}

		// Validar campos
		if nombres == "" || apellidos == "" || email == "" || password == "" {
			c.JSON(http.StatusExpectationFailed, nil)
			return
		}

		user := models.Usuario{
			Nombre:   nombres,
			Apellido: apellidos,
			Email:    email,
			Password: strconv.Quote(string(hashPassword)),
		}

		// Llamar a la función InsertNewUser de Query para insertar el nuevo usuario en la base de datos
		_, err1 := Query.InsertNewUser(user)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar el usuario"})
			return
		}

		// Devolver una respuesta al cliente
		c.JSON(http.StatusOK, gin.H{"message": "Usuario insertado exitosamente"})
	}
}
