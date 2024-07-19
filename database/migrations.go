package database

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
	// &models.Post{},
	//&models.User{}, &models.Role{}, &models.Level{},
	)
	if err != nil {
		return err
	}
	return nil
}
