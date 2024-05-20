package postgres

import (
	"github.com/omo-nosa/repository-pattern/internal"
)

type Todo struct {
	ID          int
	Title       string
	Description string
	Status      string
}

var todos internal.Todo

func NewTodo() *Todo {
	return &Todo{}
}

func (t *Todo) GetAll() ([]internal.Todo, error) {
	return []internal.Todo{
		todos,
	}, nil
}

func (t *Todo) Create(todo internal.Todo) error {
	todos = todo
	return nil
}

func (t *Todo) FindById(id int) (internal.Todo, error) {
	return internal.Todo{}, nil
}

func (t *Todo) Update(todo internal.Todo) error {
	return nil
}
