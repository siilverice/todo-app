package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siilverice/todo-app/Models"
)

func GetAllTodos(c *gin.Context) {
	var todo []Models.Todo
	if err := Models.GetAllTodos(&todo); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateTodo(c *gin.Context) {
	var todo Models.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else if err := Models.CreateTodo(&todo); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusCreated, todo)
	}
}

func GetTodo(c *gin.Context) {
	var todo Models.Todo
	c.JSON(http.StatusCreated, todo)
}
