package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"rest-http-server/db"
	"rest-http-server/models"

	"github.com/gorilla/mux"
)

func GetTodoList(w http.ResponseWriter, r *http.Request) {
	todoList := db.FindAll()

	bytes, err := json.Marshal(todoList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(w, bytes)
}

func AddTodoItem(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	todoItem := new(models.TodoItem)
	err = json.Unmarshal(body, todoItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db.Save(todoItem.ID.String(), todoItem)

	w.Header().Set("Location", r.URL.Path+"/"+todoItem.ID.String())
	w.WriteHeader(http.StatusCreated)
}

func UpdateTodoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoItemID := vars["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	todoItem := new(models.TodoItem)
	err = json.Unmarshal(body, todoItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db.Save(todoItemID, todoItem)
}

func DeleteTodoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoItemID := vars["id"]

	db.Remove(todoItemID)
	w.WriteHeader(http.StatusNoContent)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
