//main.go

package main

import (
	"fmt"
	"github.com/omo-nosa/repository-pattern/internal"
	"github.com/omo-nosa/repository-pattern/internal/postgres"
	"github.com/omo-nosa/repository-pattern/internal/service"
)

func main() {
	repo := postgres.NewTodo()
	todoService := service.NewTodo(repo)
	err := todoService.Create(internal.Todo{
		Title:       "Learn Go",
		Description: "Learn Go programming language",
		Status:      "active",
	})

	if err != nil {
		fmt.Println(err)
	}

	todos, err := todoService.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(todos)

	todo, err := todoService.FindById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(todo)

	err = todoService.Update(internal.Todo{
		ID:          1,
		Title:       "Learn Go",
		Description: "Learn Go programming language",
		Status:      "completed",
	})

	if err != nil {
		fmt.Println(err)
	}
}
