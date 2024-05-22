package get_activity_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ApusFormatJson struct {
	ServiceUsed   string
	ID            string
	Descripcion   string
	Unidad        string
	PrecioBase    float64
	FechaCreacion time.Time
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
			FechaCreacion: activity.CreatedAt,
			Insumos:       apu,
		}
		c.JSON(http.StatusOK, response)

	}
}
