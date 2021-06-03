package Models

import (
	"fmt"

	"github.com/siilverice/todo-app/Config"
)

func GetAllTodos(todo *[]Todo) (err error) {
	if err := Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *Todo) (err error) {
	if err := Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodo(todo *Todo, id string) (err error) {
	if err := Config.DB.Find(todo, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodos(todo *Todo) (err error) {
	fmt.Println(todo)
	Config.DB.Save(todo)
	return nil
}

func DeleteTodos(todo *Todo) (err error) {
	Config.DB.Delete(todo)
	return nil
}