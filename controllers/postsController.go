package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tahmid56/go-crud/initializers"
	"github.com/tahmid56/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	fmt.Println(c.Params)
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	// user := c.MustGet("userid").(string)
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	err := initializers.DB.First(&post, id)

	if err.Error != nil {
		c.JSON(404, gin.H{
			"message": fmt.Sprintf("post with id:%s doesn't exist", id),
		})
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	err := initializers.DB.First(&post)
	if err.Error != nil {
		c.JSON(404, gin.H{
			"message": fmt.Sprintf("post with id: %s doesn't exist!", id),
		})
		return
	}
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})
	c.JSON(201, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	err := initializers.DB.Delete(&models.Post{}, id)
	if err.Error != nil {
		c.JSON(404, gin.H{
			"message": fmt.Sprintf("post with id: %s doesn't exist!", id),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("post with id: %s deleted.", id),
	})
}
