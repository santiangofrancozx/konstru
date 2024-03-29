package Query

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func QueryHandler(DSN string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("codigo")
		item := c.Query("tipoItem")

		switch item {
		case "insumo":
			handleInsumoQuery(c, DSN, id)
		case "actividad":
			handleActividadQuery(c, DSN, id)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item type"})
		}
	}
}

func handleInsumoQuery(c *gin.Context, DSN, id string) {
	_, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		handleInsumoByID(c, DSN, id)
	} else {
		handleInsumoByNombre(c, DSN, id)
	}
}

func handleActividadQuery(c *gin.Context, DSN, id string) {
	_, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		handleActividadByID(c, DSN, id)
	} else {
		handleActividadByNombre(c, DSN, id)
	}
}

func handleInsumoByID(c *gin.Context, DSN, id string) {
	insumo, err := SelectInsumoByID(DSN, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, insumo)
}

func handleInsumoByNombre(c *gin.Context, DSN, id string) {
	insumos, err := SelectInsumoByNombre(DSN, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, insumos)
}

func handleActividadByID(c *gin.Context, DSN, id string) {
	actividad, err := SelectActivityByID(DSN, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actividad)
}

func handleActividadByNombre(c *gin.Context, DSN, id string) {
	actividades, err := SelectActividadByNombre(DSN, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actividades)
}
