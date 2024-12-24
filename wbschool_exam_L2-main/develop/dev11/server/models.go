package server

import "time"

// Event представляет событие в календаре
type Event struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}
