package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Subhadip006/rest/pkg/mocks"
	"github.com/gorilla/mux"
)

// GetBook handles GET /books/{id}
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	for _, book := range mocks.Books {
		if book.ID == id {
			// Pretty JSON output
			response, err := json.MarshalIndent(book, "", "  ")
			if err != nil {
				http.Error(w, "Error encoding book", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}
