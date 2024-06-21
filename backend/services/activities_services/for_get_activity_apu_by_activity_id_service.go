package activities_services

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
	CreatedBy     *string
	Insumos       []Queries.APUs
}

func createResponse(id string, descripcion string, unidad string, precioBase float64, createdAt time.Time, CreatedBy *string, apu []Queries.APUs) ApusFormatJson {
	return ApusFormatJson{
		ServiceUsed:   "GetActivityApuByActivityIdService",
		ID:            id,
		Descripcion:   descripcion,
		Unidad:        unidad,
		PrecioBase:    precioBase,
		FechaCreacion: createdAt,
		CreatedBy:     CreatedBy,
		Insumos:       apu,
	}
}

func GetActivityApuByActiityIdService() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response ApusFormatJson
		ID := c.Query("id")

		// Intentar obtener actividad global por ID
		activity, err := Adapters.SelectActivityByID(ID)
		//if err != nil {
		//	// Si no se encuentra, intentar obtener actividad de usuario por ID
		//	user, _ := Adapters.GetUserIdByToken(c)
		//	activityU, err := Adapters.SelectActivityUserById(user.ID, ID)
		//	if err != nil {
		//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//		return
		//	}
		//
		//	apu, err := Adapters.SelectApuUserByActivityId(ID)
		//	if err != nil {
		//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//		return
		//	}
		//
		//	response = createResponse(
		//		strconv.Itoa(activityU.ID_Actividad_Usuario),
		//		activityU.Descripcion,
		//		activityU.Unidad,
		//		activityU.PrecioBase,
		//		activityU.CreatedAt,
		//		apu,
		//	)
		//	c.JSON(http.StatusOK, response)
		//	return
		//}

		if err != nil {
			c.JSON(http.StatusNotFound, "No se encontro la actividad")
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
			activity.CreatedBy,
			apu,
		)
		c.JSON(http.StatusOK, response)
	}
}
