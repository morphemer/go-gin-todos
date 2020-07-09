package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joe41203/go-gin-todos/controllers"
)

// Routes all of routes
func Routes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/todos", controllers.GetTodos)
	}
	return r
}
