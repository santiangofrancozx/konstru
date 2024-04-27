package services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApusFormatJson struct {
	ServiceUsed   string
	ID            string
	Descripcion   string
	Unidad        string
	PrecioBase    float64
	FechaCreacion string
	Insumos       []Queries.APUs
}

func GetActivityApuByActiityIdService() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID := c.Query("id")
		activity, err := Adapters.SelectActivityByID(ID)
		apu, err := Adapters.SelectApuByActivityId(ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response := ApusFormatJson{
			ServiceUsed:   "GetActivityApuByActivityIdService",
			ID:            activity.ID,
			Descripcion:   activity.Descripcion,
			Unidad:        activity.Unidad,
			PrecioBase:    activity.PrecioBase,
			FechaCreacion: activity.FechaCreacion,
			Insumos:       apu,
		}
		c.JSON(http.StatusOK, response)

	}
}
