package main

import (
	"github.com/siilverice/todo-app/Config"
	"github.com/siilverice/todo-app/Models"
	"github.com/siilverice/todo-app/Routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	r.Run()
}
