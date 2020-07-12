package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/joe41203/go-gin-todos/config"
	"github.com/joe41203/go-gin-todos/models"
	"github.com/joe41203/go-gin-todos/routes"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(&models.Todo{})

	r := routes.Routes()
	r.LoadHTMLGlob("views/*/**")
	r.Run()
}
