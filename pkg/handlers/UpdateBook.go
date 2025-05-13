package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Subhadip006/rest/pkg/mocks"
	"github.com/Subhadip006/rest/pkg/models"
	"github.com/gorilla/mux"
)

// UpdateBook handles PUT /books/{id}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updatedBook models.Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	for index, book := range mocks.Books {
		if book.ID == id {
			mocks.Books[index].Title = updatedBook.Title
			mocks.Books[index].Author = updatedBook.Author
			mocks.Books[index].Price = updatedBook.Price

			response, err := json.MarshalIndent(mocks.Books, "", "  ")
			if err != nil {
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}
