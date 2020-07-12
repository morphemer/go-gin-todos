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

// GetATodo get a todo
func GetATodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	err := models.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
