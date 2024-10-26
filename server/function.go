package server

import (
	"encoding/json"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string `json:"message"`
	}{
		Message: "Hello, World!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
