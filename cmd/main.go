package main

import (
	"net/http"

	"github.com/Subhadip006/rest/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books", handlers.AddBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
