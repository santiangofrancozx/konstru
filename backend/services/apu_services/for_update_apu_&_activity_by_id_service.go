package apu_services

import (
	"awesomeKonstru/backend/handlers/Adapters"
	"awesomeKonstru/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Apu2 struct {
	InsumoID string  `json:"insumo_id"`
	Cantidad float64 `json:"cantidad"`
}

type InsumoApu2 struct {
	ID                  string  `json:"ID"`
	ActivityDescripcion string  `json:"ActivityDescripcion"`
	ActivityUnit        string  `json:"ActivityUnit"`
	ActivityBasePrice   float64 `json:"ActivityBasePrice"`
	APU                 []Apu   `json:"APU"`
}

func UpdateApuById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var InsumoApu2 InsumoApu2
		if err := c.BindJSON(&InsumoApu2); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := Adapters.GetUserIdByToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		ID := user.ID

		Activitypload := models.Actividad{
			ID:          InsumoApu2.ID,
			Descripcion: InsumoApu2.ActivityDescripcion,
			Unidad:      InsumoApu2.ActivityUnit,
			PrecioBase:  InsumoApu2.ActivityBasePrice,
			CreatedBy:   &ID,
		}
		errAU := Adapters.UpdateActivityAdapter(Activitypload)
		if errAU != nil {
			c.JSON(http.StatusInternalServerError, "Error al actualizar Actividad, eres el creador?")
			return
		}
		var uploadsNewRelations []models.ActividadInsumo
		for _, relation := range InsumoApu2.APU {
			insumosUpdate := models.ActividadInsumo{
				ActividadID: InsumoApu2.ID,
				InsumoID:    relation.InsumoID,
				Cantidad:    relation.Cantidad,
			}
			uploadsNewRelations = append(uploadsNewRelations, insumosUpdate)
		}

		errUInsimos := Adapters.UpdateActivityInsumosAdapter(InsumoApu2.ID, uploadsNewRelations)
		if errUInsimos != nil {
			c.JSON(http.StatusInternalServerError, "error al actualizar nuevos insumos")
		}

	}
}
