package server

import (
	"log"
	"net/http"
)

// StartServer запускает HTTP-сервер
func StartServer() {
	// Инициализация хранилища
	storage = NewEventStorage()

	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", CreateEventHandler)
	mux.HandleFunc("/update_event", UpdateEventHandler)
	mux.HandleFunc("/delete_event", DeleteEventHandler)

	handler := LoggingMiddleware(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Println("Server started on :8080")
	log.Fatal(server.ListenAndServe())
}
