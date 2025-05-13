package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Subhadip006/rest/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	for i, book := range mocks.Books {
		if book.ID == id {
			mocks.Books = append(mocks.Books[:i], mocks.Books[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("deleted")
			return
		}
	}
}
