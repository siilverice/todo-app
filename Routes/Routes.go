package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/siilverice/todo-app/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", Controllers.GetAllTodos)
		v1.POST("todo", Controllers.CreateTodo)
		v1.GET("todo/:id", Controllers.GetTodo)
		// v1.PUT("todo/:id", Controllers.UpdateTodos)
		// v1.DELETE("todo/:id", Controllers.DeleteTodos)
	}
	return r
}
