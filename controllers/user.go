package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	db := config.GetDB()

	if db.NewRecord(user) {
		db.Create(&user)
		c.JSON(http.StatusCreated, user)
	} else {
		c.JSON(http.StatusInternalServerError, user)
	}
}

// DeleteUser delete user
func DeleteUser(c *gin.Context) {
	id64, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	userID := uint(id64)
	user := models.User{Model: gorm.Model{ID: userID}}
	db := config.GetDB()
	if err := db.Delete(&user).Error; err != nil {
		fmt.Printf("%v\n", err)
		c.JSON(http.StatusNotAcceptable, err)
		return
	}
	c.JSON(http.StatusAccepted, user)
}
