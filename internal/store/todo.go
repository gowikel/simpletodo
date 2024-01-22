package store

import (
	"fmt"

	"github.com/gowikel/simpletodo/internal/entity"
	"github.com/gowikel/simpletodo/internal/repository"
)

type inMemoryStore struct {
	lst []entity.Todo
}

func NewInMemoryTodoStore() repository.TodoRepository {
	return &inMemoryStore{
		lst: []entity.Todo{},
	}
}

func (s inMemoryStore) FindAll() ([]entity.Todo, error) {
	return s.lst, nil
}

func (s inMemoryStore) Find(id int) (entity.Todo, error) {
	for _, todo := range s.lst {
		if todo.ID == id {
			return todo, nil
		}
	}

	return entity.Todo{}, fmt.Errorf("not found")
}

func (s *inMemoryStore) Save(todo entity.Todo) error {

	for _, savedTodo := range s.lst {
		if todo.ID == savedTodo.ID {
			return fmt.Errorf("todo already exists")
		}
	}

	s.lst = append(s.lst, todo)

	return nil
}

func (s *inMemoryStore) Update(todo entity.Todo) error {
	for i := 0; i < len(s.lst); i++ {
		if s.lst[i].ID != todo.ID {
			continue
		}

		s.lst[i] = todo
		return nil
	}

	return fmt.Errorf("todo not found")
}

func (s *inMemoryStore) Delete(id int) error {
	return nil
}
