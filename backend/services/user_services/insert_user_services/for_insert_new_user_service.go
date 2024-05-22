package insert_user_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type UsuarioRequest struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func InsertUserService() gin.HandlerFunc {
	return func(c *gin.Context) {
		var usuarioRequest UsuarioRequest
		if err := c.BindJSON(&usuarioRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(usuarioRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Error al codificar la contraseña:", err)
		}
		newUser := models.Usuario{
			Nombre:   usuarioRequest.Nombre,
			Apellido: usuarioRequest.Apellido,
			Email:    usuarioRequest.Email,
			Password: string(hashPassword),
		}

		if newUser.Nombre == "" || newUser.Apellido == "" || newUser.Email == "" || newUser.Password == "" {
			c.JSON(http.StatusExpectationFailed, nil)
			return
		}

		// Llamar a la función InsertNewUser de Query para insertar el nuevo usuario en la base de datos
		_, err1 := Adapters.InsertNewUser(newUser)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar el usuario"})
			return
		}

		// Devolver una respuesta al cliente
		c.JSON(http.StatusOK, gin.H{"message": "Usuario insertado exitosamente"})
	}
}
