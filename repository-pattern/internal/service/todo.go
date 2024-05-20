package service

import (
	"github.com/omo-nosa/repository-pattern/internal"
	
)

type TodoRepository interface {
	GetAll() ([]internal.Todo, error)
	Create(todo internal.Todo) error
	FindById(id int) (internal.Todo, error)
    Update(todo internal.Todo) error
}


type todo struct {
	repo TodoRepository
}

func NewTodo(repo TodoRepository) *todo {
	return &todo{repo} // returns a pointer to the todo struct
}

func (t *todo) GetAll() ([]internal.Todo, error) {
	return t.repo.GetAll()
}

func (t *todo) Create(todo internal.Todo) error {
	return t.repo.Create(todo)
}

func (t *todo) FindById(id int) (internal.Todo, error) {
	return t.repo.FindById(id)
}

func (t *todo) Update(todo internal.Todo) error {
	return t.repo.Update(todo)
}
