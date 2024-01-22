package main

import (
	"net/http"

	"github.com/gorilla/mux"
	httpHandler "github.com/gowikel/simpletodo/internal/delivery/http"
	"github.com/gowikel/simpletodo/internal/log"
	"github.com/gowikel/simpletodo/internal/service"
	store "github.com/gowikel/simpletodo/internal/store/inMemory"
)

func main() {
	store := store.NewInMemoryTodoStore()
	todoService := service.NewTodoService(store)
	todoHandler := httpHandler.NewTodoHandler(todoService)

	r := mux.NewRouter()

	r.HandleFunc("/todos", todoHandler.GetAllTodos).Methods("GET")
	r.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", todoHandler.GetTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).
		Methods("POST", "PUT")
	r.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).
		Methods("DELETE")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
