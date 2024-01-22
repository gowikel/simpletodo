package repository

import (
	"github.com/gowikel/simpletodo/internal/entity"
)

type TodoFilters struct {
	Completed bool
}

type TodoReader interface {
	FindAll() ([]entity.Todo, error)
	Find(id int) (entity.Todo, error)
	Query(filters TodoFilters) ([]entity.Todo, error)
}

type TodoWriter interface {
	Save(todo entity.Todo) error
	Delete(id int) error
	Update(todo entity.Todo) error
}

type TodoRepository interface {
	TodoReader
	TodoWriter
}
