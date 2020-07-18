package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joe41203/go-gin-todos/config"
	"github.com/joe41203/go-gin-todos/models"
)

// GetAllUsers fetch all users
func GetAllUsers(c *gin.Context) {
	users := []models.User{}
	db := config.GetDB()
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// CreateUser create  a user
func CreateUser(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)

	if err := user.Validate(); err != nil {
		fmt.Printf("%v\n", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	db := config.GetDB()

	if db.NewRecord(user) {
		db.Create(&user)
		c.JSON(http.StatusCreated, user)
	} else {
		c.JSON(http.StatusInternalServerError, user)
	}
}
