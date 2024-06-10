package activities_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/handlers/Adapters/Queries"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func createResponse(id string, descripcion string, unidad string, precioBase float64, createdAt time.Time, apu []Queries.APUs) ApusFormatJson {
	return ApusFormatJson{
		ServiceUsed:   "GetActivityApuByActivityIdService",
		ID:            id,
		Descripcion:   descripcion,
		Unidad:        unidad,
		PrecioBase:    precioBase,
		FechaCreacion: createdAt,
		Insumos:       apu,
	}
}

func GetActivityApuByActiityIdService() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response ApusFormatJson
		ID := c.Query("id")

		// Intentar obtener actividad global por ID
		activity, err := Adapters.SelectActivityByID(ID)
		if err != nil {
			// Si no se encuentra, intentar obtener actividad de usuario por ID
			user, _ := Adapters.GetUserIdByToken(c)
			activityU, err := Adapters.SelectActivityUserById(user.ID, ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			apu, err := Adapters.SelectApuUserByActivityId(ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			response = createResponse(
				strconv.Itoa(activityU.ID_Actividad_Usuario),
				activityU.Descripcion,
				activityU.Unidad,
				activityU.PrecioBase,
				activityU.CreatedAt,
				apu,
			)
			c.JSON(http.StatusOK, response)
			return
		}

		// Si se encuentra la actividad global
		apu, err := Adapters.SelectApuByActivityId(ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response = createResponse(
			activity.ID,
			activity.Descripcion,
			activity.Unidad,
			activity.PrecioBase,
			activity.CreatedAt,
			apu,
		)
		c.JSON(http.StatusOK, response)
	}
}
