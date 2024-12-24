package server

import (
	"encoding/json"
	"net/http"
)

// WriteJSON отправляет JSON-ответ
func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
