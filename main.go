package main

import (
	"log"

	"github.com/uxioandrade/go-sql-tutorial/model/todos"
)

func main() {
	todo := todos.Todo{
		Description: "First todo",
		Priority:    1,
		Status:      "In Progress",
	}
	todo2 := todos.Todo{
		Description: "Second todo",
		Priority:    3,
		Status:      "Done",
	}
	oldTodo := todos.Todo{
		ID: 1,
	}
	todo.Save()
	todo2.Save()
	oldTodo.Get()
	log.Println(oldTodo)
}
