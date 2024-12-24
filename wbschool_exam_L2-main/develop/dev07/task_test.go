package main

import (
	"testing"
	"time"
)

// вспомогательная функция для создания канала с задержкой
func sigTest(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

// Тестирование функции or
func TestOr(t *testing.T) {
	t.Run("No channels", func(t *testing.T) {
		done := or()
		select {
		case <-done:
			// Ожидаем немедленное закрытие
		case <-time.After(100 * time.Millisecond):
			t.Fatal("Expected closed channel immediately with no channels")
		}
	})

	t.Run("Single channel", func(t *testing.T) {
		done := or(sigTest(100 * time.Millisecond))
		select {
		case <-done:
			// Ожидаем закрытие после 100ms
		case <-time.After(200 * time.Millisecond):
			t.Fatal("Expected channel to close after 100ms")
		}
	})

	t.Run("Two channels", func(t *testing.T) {
		done := or(
			sigTest(200*time.Millisecond),
			sigTest(100*time.Millisecond),
		)
		start := time.Now()
		<-done
		duration := time.Since(start)
		if duration > 150*time.Millisecond {
			t.Fatalf("Expected channel to close after ~100ms, took %v", duration)
		}
	})

	t.Run("Multiple channels", func(t *testing.T) {
		done := or(
			sigTest(1*time.Second),
			sigTest(500*time.Millisecond),
			sigTest(100*time.Millisecond),
		)
		start := time.Now()
		<-done
		duration := time.Since(start)
		if duration > 150*time.Millisecond {
			t.Fatalf("Expected channel to close after ~100ms, took %v", duration)
		}
	})
}
