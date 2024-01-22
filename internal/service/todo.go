package service

import (
	"fmt"

	"github.com/gowikel/simpletodo/internal/entity"
	"github.com/gowikel/simpletodo/internal/repository"
)

type TodoService interface {
	GetAllTodos() ([]entity.Todo, error)
	GetUncompletedTodos() ([]entity.Todo, error)
	GetTodo(id int) (entity.Todo, error)
	CreateTodo(todo entity.Todo) error
	UpdateTodo(todo entity.Todo) error
	DeleteTodo(id int) error
}

type todoServiceImpl struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return todoServiceImpl{repo: repo}
}

func (s todoServiceImpl) GetAllTodos() ([]entity.Todo, error) {
	result, err := s.repo.FindAll()

	if err != nil {
		return result, fmt.Errorf("TodoService: %w", err)
	}

	return result, nil
}

func (s todoServiceImpl) GetUncompletedTodos() ([]entity.Todo, error) {
	result, err := s.repo.Query(repository.TodoFilters{
		Completed: false,
	})

	if err != nil {
		return result, fmt.Errorf("TodoService: %w", err)
	}

	return result, nil
}

func (s todoServiceImpl) GetTodo(id int) (entity.Todo, error) {
	result, err := s.repo.Find(id)

	if err != nil {
		return result, fmt.Errorf("TodoService: %w", err)
	}

	return result, nil
}

func (s todoServiceImpl) CreateTodo(todo entity.Todo) error {
	err := s.repo.Save(todo)

	if err != nil {
		return fmt.Errorf("TodoService: %w", err)
	}

	return nil
}

func (s todoServiceImpl) UpdateTodo(todo entity.Todo) error {
	err := s.repo.Update(todo)

	if err != nil {
		return fmt.Errorf("TodoService: %w", err)
	}

	return nil
}

func (s todoServiceImpl) DeleteTodo(id int) error {
	err := s.repo.Delete(id)

	if err != nil {
		return fmt.Errorf("TodoService: %w", err)
	}

	return nil
}
