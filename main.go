package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", healthHandler)
	http.ListenAndServe(":5000", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"response": "ok"}
	json.NewEncoder(w).Encode(response)
}

// Комментарий ещё