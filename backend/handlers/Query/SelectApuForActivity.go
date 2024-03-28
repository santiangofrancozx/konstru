package Query

import (
	"awesomeKonstru/backend/handlers"
	"awesomeKonstru/backend/models"
	"gorm.io/gorm"
)

type Apu struct {
	Actividad models.Actividad
	Insumos   []models.Insumo
}

func QueryInsumoActivityByActivityID(db *gorm.DB, activity_id string) ([]models.ActividadInsumo, error) {
	var items []models.ActividadInsumo
	if err := db.Where("actividad_id = ?", activity_id).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func SelectInsumoActivityByActivityID(dsn string, activity_id string) ([]models.ActividadInsumo, error) {
	db, err := handlers.Connect(dsn)
	if err != nil {
		return nil, err
	}
	query, err := QueryInsumoActivityByActivityID(db, activity_id)
	defer handlers.Disconnect(db)

	return query, err
}

func QueryInsumoByListInsumoActivity(db *gorm.DB, insumos []models.ActividadInsumo) ([]models.Insumo, error) {
	var IL []models.Insumo
	for _, insumo := range insumos {
		i1, err := QueryInsumoByID(db, insumo.InsumoID)
		if err != nil {
			return nil, err
		}
		IL = append(IL, i1)
	}
	return IL, nil
}
func SelectInsumosByListActivityInsumo(dsn string, insumos []models.ActividadInsumo) ([]models.Insumo, error) {
	db, err := handlers.Connect(dsn)
	if err != nil {
		return nil, err
	}
	query, err := QueryInsumoByListInsumoActivity(db, insumos)
	defer handlers.Disconnect(db)

	return query, err
}
