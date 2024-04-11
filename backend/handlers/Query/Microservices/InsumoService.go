package Microservices

import (
	"awesomeKonstru/backend/handlers/Query"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleInsumoQuery(c *gin.Context, DSN, id string) {
	_, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		handleInsumoByID(c, DSN, id)
	} else {
		handleInsumoByNombre(c, DSN, id)
	}
}

func handleInsumoByID(c *gin.Context, DSN, id string) {
	insumo, err := Query.SelectInsumoByID(DSN, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, insumo)
}

func handleInsumoByNombre(c *gin.Context, DSN, id string) {
	insumos, err := Query.SelectInsumoByNombre(DSN, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, insumos)
}
