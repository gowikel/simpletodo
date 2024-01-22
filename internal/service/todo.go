package service

import (
	"fmt"

	"github.com/gowikel/simpletodo/internal/entity"
	"github.com/gowikel/simpletodo/internal/repository"
)

type TodoService interface {
	GetAllTodos() ([]entity.Todo, error)
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

func (r todoServiceImpl) GetAllTodos() ([]entity.Todo, error) {
	result, err := r.repo.FindAll()

	if err != nil {
		return result, fmt.Errorf("TodoService: %w", err)
	}

	return result, nil
}

func (r todoServiceImpl) GetTodo(id int) (entity.Todo, error) {
	result, err := r.repo.Find(id)

	if err != nil {
		return result, fmt.Errorf("TodoService: %w", err)
	}

	return result, nil
}

func (r todoServiceImpl) CreateTodo(todo entity.Todo) error {
	err := r.repo.Save(todo)

	if err != nil {
		return fmt.Errorf("TodoService: %w", err)
	}

	return nil
}

func (r todoServiceImpl) UpdateTodo(todo entity.Todo) error {
	err := r.repo.Update(todo)

	if err != nil {
		return fmt.Errorf("TodoService: %w", err)
	}

	return nil
}

func (r todoServiceImpl) DeleteTodo(id int) error {
	err := r.repo.Delete(id)

	if err != nil {
		return fmt.Errorf("TodoService: %w", err)
	}

	return nil
}
