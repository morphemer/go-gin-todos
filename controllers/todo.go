// Package controllers todos
package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetTodos get all todos
func GetTodos(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello"})
}
