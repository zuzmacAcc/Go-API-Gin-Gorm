package main

import (
	"github.com/zuzmacAcc/Go-API-Gin-Gorm/initializers"
	"github.com/zuzmacAcc/Go-API-Gin-Gorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	 initializers.DB.AutoMigrate(&models.Post{})
	
}
