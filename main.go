package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joe41203/go-gin-todos/config"
	"github.com/joe41203/go-gin-todos/controllers"
)

func main() {
	fmt.Println("HelloWorld")
	db := config.Init()
	defer db.Close()

	r := gin.Default()

	r.GET("/", controllers.GetAllUsers)

	r.Run()
}
