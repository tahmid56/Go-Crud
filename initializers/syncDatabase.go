package initializers

import "github.com/tahmid56/go-crud/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}