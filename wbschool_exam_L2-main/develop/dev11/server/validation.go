package server

import (
	"errors"
	"net/url"
	"time"
)

// ValidateEventParams проверяет параметры события
func ValidateEventParams(params url.Values) (Event, error) {
	id := params.Get("id")
	userID := params.Get("user_id")
	title := params.Get("title")
	dateStr := params.Get("date")

	if id == "" || userID == "" || title == "" || dateStr == "" {
		return Event{}, errors.New("missing required parameters")
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return Event{}, errors.New("invalid date format")
	}

	return Event{
		ID:        id,
		UserID:    userID,
		Title:     title,
		Date:      date,
		CreatedAt: time.Now(),
	}, nil
}
