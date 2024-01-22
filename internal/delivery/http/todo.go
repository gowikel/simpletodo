package http

import (
	"encoding/json"
	"net/http"

	"github.com/gowikel/simpletodo/internal/service"
)

type TodoHandler struct {
	TodoService service.TodoService
}

func NewTodoHandler(service service.TodoService) TodoHandler {
	return TodoHandler{
		TodoService: service,
	}
}

func (h TodoHandler) GetAllTodos(
	w http.ResponseWriter,
	r *http.Request,
) {
	todos, err := h.TodoService.GetAllTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}
