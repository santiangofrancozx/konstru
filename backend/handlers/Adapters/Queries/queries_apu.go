package Queries

import (
	"awesomeKonstru/backend/models"
	"gorm.io/gorm"
	"time"
)

type APUs struct {
	ID                 string
	Descripcion        string
	Unidad             string
	PrecioBase         float64
	FechaActualizacion time.Time
	Clasificacion      string
	Cantidad           float64
}

func QueryAPUByActivityID(db *gorm.DB, id string) ([]APUs, error) {
	var insumos []APUs
	if err := db.
		Model(&models.ActividadInsumo{}).
		Select("insumos.*, actividad_insumos.cantidad").
		Joins("JOIN insumos ON actividad_insumos.insumo_id = insumos.id").
		Where("actividad_insumos.actividad_id = ?", id).
		Find(&insumos).Error; err != nil {
		return nil, err
	}
	return insumos, nil
}
func QueryAPUUserByActivityID(db *gorm.DB, id string) ([]APUs, error) {
	var insumos []APUs
	if err := db.
		Model(&models.ActividadU_InsumoU{}).
		Select("insumos.*, actividadu_insumo_us.cantidad").
		Joins("JOIN insumos ON actividadu_insumo_us.insumo_uid = insumos.id").
		Where("actividadu_insumo_us.actividad_uid = ?", id).
		Find(&insumos).Error; err != nil {
		return nil, err
	}
	return insumos, nil
}

func QueryInsertNewAPUs(db *gorm.DB, apu []models.ActividadInsumo) error {
	if err := db.Create(&apu).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAPUByID(db *gorm.DB, id string) error {
	// Iniciar una transacción
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error // Manejar el error si no se puede iniciar la transacción
	}

	// Deferir la función para hacer rollback si hay un error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Hacer rollback si hay un error panicking
		} else if r != nil {
			tx.Rollback() // Hacer rollback si hay un error normal
		} else {
			tx.Commit() // Confirmar la transacción si no hay errores
		}
	}()

	// Eliminar los registros relacionados
	if err := tx.Where("actividad_id = ?", id).Delete(&models.ActividadInsumo{}).Error; err != nil {
		return err // Retornar el error si falla la eliminación
	}

	return nil // Retornar nil si todo se realiza correctamente
}

func QueryInsertNewAPUsUser(db *gorm.DB, apu []models.ActividadU_InsumoU) error {
	if err := db.Create(&apu).Error; err != nil {
		return err
	}
	return nil
}
