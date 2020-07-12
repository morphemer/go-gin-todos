// Package controllers todos
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joe41203/go-gin-todos/models"
)

// GetTodos get all todos
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	err := models.GetAllTodos(&todos)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todos)
	}
}
