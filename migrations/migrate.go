package main

import (
	"github.com/tahmid56/go-crud/initializers"
	"github.com/tahmid56/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
