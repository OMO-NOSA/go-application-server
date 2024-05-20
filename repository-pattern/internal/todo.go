package internal

import "errors"


var (
	ErrTitleCannotBeEmpty = errors.New("title cannot be empty")
	ErrDescriptionCannotBeEmpty = errors.New("description cannot be empty")
)

type Todo struct {
	ID          int
	Title       string
	Description string
	Status	    string

}

func (t *Todo) Validate() error {
	if t.Title == "" {
		return ErrTitleCannotBeEmpty 
	}
	
	if t.Description == "" {
		return ErrDescriptionCannotBeEmpty
	}

	return nil
}
