package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tahmid56/go-crud/initializers"
	"github.com/tahmid56/go-crud/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context){
	var body struct {
		Name string
		Email string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body!",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash password",
		})
		return
	}
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to create user!",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully!",
	})
}