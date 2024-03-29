package Query

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryHandler(DSN string) gin.HandlerFunc {
	// Suponiendo que el ID del insumo proviene de la URL como un parámetr
	return func(c *gin.Context) {
		// Suponiendo que el ID del insumo proviene de la URL como un parámetro
		id := c.Query("codigo")

		fmt.Println(id)

		// Realizar la consulta del insumo por su ID utilizando la lógica definida en el paquete Query
		insumo, err := SelectInsumoByID(DSN, id)
		if err != nil {
			// Manejar el error devolviendo un mensaje de error JSON
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Devolver los datos del insumo como JSON
		c.JSON(http.StatusOK, insumo)
	}
}
