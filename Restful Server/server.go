package main

import (
	"net/http"
	"rest-http-server/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Methods("GET").Path("/todo-list").HandlerFunc(handlers.GetTodoList)
	router.Methods("POST").Path("/todo-list").HandlerFunc(handlers.AddTodoItem)
	router.Methods("PUT").Path("/todo-list/{id}").HandlerFunc(handlers.UpdateTodoItem)
	router.Methods("DELETE").Path("/todo-list/{id}").HandlerFunc(handlers.DeleteTodoItem)

	http.ListenAndServe(":8080", router)
}
