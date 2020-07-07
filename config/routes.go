package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/todos", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello"})
		})
	}
	return r
}
