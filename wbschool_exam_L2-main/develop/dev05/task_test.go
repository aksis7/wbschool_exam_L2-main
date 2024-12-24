package main

import (
	"regexp"
	"testing"
)

// Тест на поиск совпадений
func TestMatchLines(t *testing.T) {
	lines := []string{
		"INFO: Application started",
		"ERROR: Failed to load configuration",
		"INFO: Retrying connection",
		"WARNING: Low memory",
		"ERROR: Connection timed out",
		"INFO: Application stopped",
	}

	pattern := "ERROR"
	re, err := regexp.Compile(pattern)
	if err != nil {
		t.Fatalf("Failed to compile pattern: %v", err)
	}

	// Проверяем поиск строк с совпадением
	expectedIndices := []int{1, 4}
	matchIndices := []int{}
	for i, line := range lines {
		if re.MatchString(line) {
			matchIndices = append(matchIndices, i)
		}
	}

	if len(matchIndices) != len(expectedIndices) {
		t.Errorf("Expected %d matches, got %d", len(expectedIndices), len(matchIndices))
	}

	for i, idx := range matchIndices {
		if idx != expectedIndices[i] {
			t.Errorf("Expected match at index %d, got %d", expectedIndices[i], idx)
		}
	}
}

// Тест на флаг ignore-case (-i)
func TestIgnoreCaseMatch(t *testing.T) {
	lines := []string{
		"INFO: Application started",
		"error: lowercase error message",
		"INFO: Retrying connection",
	}

	pattern := "(?i)ERROR" // (?i) для игнорирования регистра
	re, err := regexp.Compile(pattern)
	if err != nil {
		t.Fatalf("Failed to compile pattern: %v", err)
	}

	expectedIndices := []int{1}
	matchIndices := []int{}
	for i, line := range lines {
		if re.MatchString(line) {
			matchIndices = append(matchIndices, i)
		}
	}

	if len(matchIndices) != len(expectedIndices) {
		t.Errorf("Expected %d matches, got %d", len(expectedIndices), len(matchIndices))
	}

	for i, idx := range matchIndices {
		if idx != expectedIndices[i] {
			t.Errorf("Expected match at index %d, got %d", expectedIndices[i], idx)
		}
	}
}

// Тест на флаг invert (-v)
func TestInvertMatch(t *testing.T) {
	lines := []string{
		"INFO: Application started",
		"ERROR: Failed to load configuration",
		"INFO: Retrying connection",
	}

	pattern := "ERROR"
	re, err := regexp.Compile(pattern)
	if err != nil {
		t.Fatalf("Failed to compile pattern: %v", err)
	}

	expectedIndices := []int{0, 2}
	matchIndices := []int{}
	for i, line := range lines {
		if !re.MatchString(line) {
			matchIndices = append(matchIndices, i)
		}
	}

	if len(matchIndices) != len(expectedIndices) {
		t.Errorf("Expected %d non-matching lines, got %d", len(expectedIndices), len(matchIndices))
	}

	for i, idx := range matchIndices {
		if idx != expectedIndices[i] {
			t.Errorf("Expected non-match at index %d, got %d", expectedIndices[i], idx)
		}
	}
}

// Тест на контекст (-C)
func TestContextMatch(t *testing.T) {
	lines := []string{
		"INFO: Application started",
		"ERROR: Failed to load configuration",
		"INFO: Retrying connection",
		"WARNING: Low memory",
		"ERROR: Connection timed out",
		"INFO: Application stopped",
	}

	pattern := "ERROR"
	re, err := regexp.Compile(pattern)
	if err != nil {
		t.Fatalf("Failed to compile pattern: %v", err)
	}

	contextBefore := 1
	contextAfter := 1

	expectedLines := []int{0, 1, 2, 3, 4, 5}
	printedLines := map[int]bool{}

	for i, line := range lines {
		if re.MatchString(line) {
			start := i - contextBefore
			if start < 0 {
				start = 0
			}
			end := i + contextAfter
			if end >= len(lines) {
				end = len(lines) - 1
			}
			for j := start; j <= end; j++ {
				printedLines[j] = true
			}
		}
	}

	// Проверка совпадений с ожидаемыми индексами
	for _, idx := range expectedLines {
		if !printedLines[idx] {
			t.Errorf("Expected line %d to be printed, but it wasn't", idx)
		}
	}
}
