package db

import (
	"github.com/qor/auth/auth_identity"
)

func MigrateAll() {
	AutoMigrate(&auth_identity.AuthIdentity{})
}

// AutoMigrate run auto migration
func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		DB.AutoMigrate(value)
	}
}
