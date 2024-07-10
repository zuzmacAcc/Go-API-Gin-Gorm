package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zuzmacAcc/Go-API-Gin-Gorm/initializers"
	"github.com/zuzmacAcc/Go-API-Gin-Gorm/models"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	if initializers.DB == nil {
		log.Fatal("DB is nil in CreatePost")
	}

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	if initializers.DB == nil {
		log.Fatal("DB is nil in CreatePost")
	}
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	if initializers.DB == nil {
		log.Fatal("DB is nil in CreatePost")
	}

	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}
