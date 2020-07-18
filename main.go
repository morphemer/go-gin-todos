package main

import (
	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	r := gin.Default()
	r.Use(cors.New(config))

	r.GET("/", controllers.GetAllUsers)
	r.POST("/users", controllers.CreateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.Run()
}
