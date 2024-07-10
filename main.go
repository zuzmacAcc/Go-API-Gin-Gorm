package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zuzmacAcc/Go-API-Gin-Gorm/controllers"
	"github.com/zuzmacAcc/Go-API-Gin-Gorm/initializers"
)

func init() {
	log.Println("Initializing environment variables and database connection")
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	log.Println("Starting application")
	r := gin.Default()
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/post/:id", controllers.GetPost)
	r.Run()
}
