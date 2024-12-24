package server

import (
	"errors"
	"sync"
)

// Event представляет событие

// EventStorage - структура для хранения событий
type EventStorage struct {
	mu     sync.RWMutex
	events map[string]Event
}

// NewEventStorage создает новое хранилище событий
func NewEventStorage() *EventStorage {
	return &EventStorage{
		events: make(map[string]Event),
	}
}

// CreateEvent добавляет событие в хранилище
func (s *EventStorage) CreateEvent(e Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.events[e.ID]; exists {
		return errors.New("event already exists")
	}
	s.events[e.ID] = e
	return nil
}

// UpdateEvent обновляет событие в хранилище
func (s *EventStorage) UpdateEvent(e Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.events[e.ID]; !exists {
		return errors.New("event not found")
	}
	s.events[e.ID] = e
	return nil
}

// DeleteEvent удаляет событие из хранилища
func (s *EventStorage) DeleteEvent(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.events[id]; !exists {
		return errors.New("event not found")
	}
	delete(s.events, id)
	return nil
}
