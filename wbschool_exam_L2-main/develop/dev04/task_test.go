package main

import (
	"reflect"
	"testing"
)

// Тестирование функции FindAnagramSets
func TestFindAnagramSets(t *testing.T) {
	// Тестовые данные
	words := []string{
		"пятак", "пятка", "тяпка",
		"листок", "слиток", "столик",
		"яблоко", "кот", "ток", "окт",
		"один",
	}

	expected := map[string][]string{
		"пятак":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
		"кот":    {"кот", "окт", "ток"},
	}

	// Выполнение функции
	result := FindAnagramSets(words)

	// Проверка на соответствие ожидаемым результатам
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

// Дополнительный тест с пустым входом
func TestFindAnagramSetsEmptyInput(t *testing.T) {
	words := []string{}

	expected := map[string][]string{}

	result := FindAnagramSets(words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось пустое множество, получено %v", result)
	}
}

// Дополнительный тест с дубликатами
func TestFindAnagramSetsDuplicates(t *testing.T) {
	words := []string{
		"пятак", "пятак", "ПЯТКА", "ТЯПКА",
	}

	expected := map[string][]string{
		"пятак": {"пятак", "пятка", "тяпка"},
	}

	result := FindAnagramSets(words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}
