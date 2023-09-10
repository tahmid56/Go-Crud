package main

import (
	"os"

	"github.com/gin-gonic/gin"

	// "github.com/tahmid56/go-crud/controllers"
	"github.com/tahmid56/go-crud/controllers"
	"github.com/tahmid56/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()

	r.POST("/sign-up", controllers.SignUp)

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run(os.Getenv("PORT"))
}
