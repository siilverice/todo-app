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
	// TODO: Validate required field
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
	// TODO: handle id not found, should return nil instead of default todo
	id := c.Param("id")
	var todo Models.Todo

	if err := Models.GetTodo(&todo, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo Models.Todo

	// TODO: handle id not found
	if err := Models.GetTodo(&todo, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if err := c.BindJSON(&todo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if err := Models.UpdateTodo(&todo); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
