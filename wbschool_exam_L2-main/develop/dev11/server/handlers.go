package server

import (
	"net/http"
)

// Глобальная переменная storage
var storage *EventStorage

// CreateEventHandler - обработчик для создания события
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	event, err := ValidateEventParams(r.Form)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err = storage.CreateEvent(event)
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"result": "event created"})
}

// UpdateEventHandler - обработчик для обновления события
func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	event, err := ValidateEventParams(r.Form)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err = storage.UpdateEvent(event)
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"result": "event updated"})
}

// DeleteEventHandler - обработчик для удаления события
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	if id == "" {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "missing id parameter"})
		return
	}

	err := storage.DeleteEvent(id)
	if err != nil {
		WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{"result": "event deleted"})
}
