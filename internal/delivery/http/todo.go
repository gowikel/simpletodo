package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gowikel/simpletodo/internal/entity"
	"github.com/gowikel/simpletodo/internal/service"
)

var decoder = schema.NewDecoder()

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

func (h TodoHandler) GetUncompletedTodos(
	w http.ResponseWriter,
	r *http.Request,
) {
	todos, err := h.TodoService.GetUncompletedTodos()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func (h TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := h.TodoService.GetTodo(todoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (h TodoHandler) CreateTodo(
	w http.ResponseWriter,
	r *http.Request,
) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var todo entity.Todo
	err = decoder.Decode(&todo, r.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.TodoService.CreateTodo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h TodoHandler) UpdateTodo(
	w http.ResponseWriter,
	r *http.Request,
) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.TodoService.GetTodo(todoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var todo entity.Todo
	err = decoder.Decode(&todo, r.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.TodoService.UpdateTodo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TodoHandler) DeleteTodo(
	w http.ResponseWriter,
	r *http.Request,
) {
	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.TodoService.DeleteTodo(todoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
