package apu_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Apu struct {
	InsumoID string  `json:"insumo_id"`
	Cantidad float64 `json:"cantidad"`
}

type InsumoApu struct {
	ActivityDescripcion string  `json:"ActivityDescripcion"`
	ActivityUnit        string  `json:"ActivityUnit"`
	ActivityBasePrice   float64 `json:"ActivityBasePrice"`
	APU                 []Apu   `json:"APU"`
}

func InsertNewApuAndActivityUserService() gin.HandlerFunc {
	return func(c *gin.Context) {
		var InsumoApu InsumoApu
		if err := c.BindJSON(&InsumoApu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := Adapters.GetUserIdByToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		ID := user.ID

		//Activitypload := models.Actividad_Usuario{
		//	Usuario_ID:  user.ID,
		//	Descripcion: InsumoApu.ActivityDescripcion,
		//	Unidad:      InsumoApu.ActivityUnit,
		//	PrecioBase:  InsumoApu.ActivityBasePrice,
		//}
		Activitypload := models.Actividad{
			Descripcion: InsumoApu.ActivityDescripcion,
			Unidad:      InsumoApu.ActivityUnit,
			PrecioBase:  InsumoApu.ActivityBasePrice,
			CreatedBy:   &ID,
		}
		errI, id := Adapters.InsertIntoActivityAdapter(Activitypload)
		if errI != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errI.Error()})
			return // Aquí se detiene la ejecución de la función si hay un error
		}
		var Apus []models.ActividadInsumo
		for _, value := range InsumoApu.APU {
			apu := models.ActividadInsumo{
				ActividadID: id,
				InsumoID:    value.InsumoID,
				Cantidad:    value.Cantidad,
			}
			Apus = append(Apus, apu)
		}
		errA := Adapters.InsertApuUserAdapter(Apus)

		if errA != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errA.Error()})
			return // Aquí se detiene la ejecución de la función si hay un error
		}
		c.JSON(http.StatusOK, InsumoApu)
	}
}
