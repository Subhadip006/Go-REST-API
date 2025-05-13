package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Subhadip006/rest/pkg/mocks"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(mocks.Books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
