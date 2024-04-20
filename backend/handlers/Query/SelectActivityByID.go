package Query

import (
	"awesomeKonstru/backend/handlers/Connection-Migrates"
	"awesomeKonstru/backend/models"
	"gorm.io/gorm"
)

func QueryActivityByID(db *gorm.DB, id string) (models.Actividad, error) {
	item := models.Actividad{}
	if err := db.First(&item, "id = ?", id).Error; err != nil {
		return models.Actividad{}, err
	}
	return item, nil
}
func QueryActividadByNombre(db *gorm.DB, nombre string) ([]models.Actividad, error) {
	var items []models.Actividad
	if err := db.Where("descripcion LIKE ?", "%"+nombre+"%").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func SelectActivityByID(id string) (models.Actividad, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return models.Actividad{}, err
	}
	defer Connection_Migrates.Disconnect(db)

	return QueryActivityByID(db, id)
}
func SelectActividadByNombre(name string) ([]models.Actividad, error) {
	db, err := Connection_Migrates.Connect()
	if err != nil {
		return nil, err
	}
	defer Connection_Migrates.Disconnect(db)

	return QueryActividadByNombre(db, name)
}
