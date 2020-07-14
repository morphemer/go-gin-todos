package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joe41203/go-gin-todos/config"
	"github.com/joe41203/go-gin-todos/controllers"
	"github.com/joe41203/go-gin-todos/models"
)

func main() {
	db := config.Init()
	defer db.Close()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Article{})

	r := gin.Default()

	r.GET("/", controllers.GetAllUsers)
	r.Run()
}
