package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Subhadip006/rest/pkg/mocks"
	"github.com/Subhadip006/rest/pkg/models"
)

func AddBook(w http.ResponseWriter, r *http.Request) {

	var newBook models.Book

	// Decode the incoming JSON body
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log the received book data for debugging
	fmt.Println("Received book:", newBook)

	// Append the new book to the list
	mocks.Books = append(mocks.Books, newBook)

	// Log the updated books list for debugging
	fmt.Println("Updated books list:", mocks.Books)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mocks.Books)
}
