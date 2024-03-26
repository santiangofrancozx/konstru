package handlers

import (
	"fmt"
	"gorm.io/gorm"
)

type Migration interface {
	Migrate() error
}

func MakeMigrations(db *gorm.DB, models []interface{}) error {
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("error al realizar las migraciones: %v", err)
	}
	return nil
}
