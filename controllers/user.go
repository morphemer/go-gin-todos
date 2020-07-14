package controllers

import (
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
