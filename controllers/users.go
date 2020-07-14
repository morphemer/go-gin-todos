package controllers

import "github.com/gin-gonic/gin"

// GetAllUsers fetch all users
func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{"message": "HelloWorld"})
}
