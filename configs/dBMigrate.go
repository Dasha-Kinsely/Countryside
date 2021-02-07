package configs

import (
	"Countryside/models"
)

// Register Models to be Migrated Here
func MigrateDB() {
	db := GetDB()
	db.AutoMigrate(&models.UserModel{})
}
