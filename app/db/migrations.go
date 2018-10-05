package db

import (
	"goqor1.0/app/models"
)

func MigrateAll() {
	AutoMigrate(&models.User{}, &models.UserAuthIdentity{})
}

// AutoMigrate run auto migration
func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		DB.AutoMigrate(value)
	}
}
