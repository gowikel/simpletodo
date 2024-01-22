package main

import (
	"log"
	"net/http"

	httpHandler "github.com/gowikel/simpletodo/internal/delivery/http"
	"github.com/gowikel/simpletodo/internal/service"
	"github.com/gowikel/simpletodo/internal/store"
)

func main() {
	store := store.NewInMemoryTodoStore()
	todoService := service.NewTodoService(store)
	todoHandler := httpHandler.NewTodoHandler(todoService)

	http.HandleFunc("/todos", todoHandler.GetAllTodos)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
