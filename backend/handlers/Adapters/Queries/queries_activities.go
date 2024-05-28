package Queries

import (
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
func InsertInToActivitiesUser(db *gorm.DB, activityUser models.Actividad_Usuario) (error, int) {
	var lastInsertedID int
	if err := db.Create(&activityUser).Error; err != nil {
		return err, -1
	}
	var activityUser2 models.Actividad_Usuario
	if err := db.Last(&activityUser2).Error; err != nil {
		// Manejar el error aquí
		return err, -1
	} else {
		lastInsertedID = activityUser2.ID_Actividad_Usuario
		// Utilizar lastInsertedID
	}
	return nil, lastInsertedID
}
func InsertInToActivities(db *gorm.DB, activity models.Actividad) error {
	if err := db.Create(&activity).Error; err != nil {
		return err
	}
	return nil
}
