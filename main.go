package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joe41203/go-gin-todos/config"
)

func main() {
	fmt.Println("HelloWorld")
	db := config.Init()
	defer db.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "HelloWorld"})
	})

	r.Run()
}
