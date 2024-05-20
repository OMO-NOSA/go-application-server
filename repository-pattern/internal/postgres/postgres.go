package postgres

type Todo struct {
	ID		  int
	Title	  string
	Description string
	Status	  string
}

func NewTodo() *Todo {
	return &Todo{}
}
