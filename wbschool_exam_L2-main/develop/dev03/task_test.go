package main

import (
	"sort"
	"testing"
)

// TestMonthSort проверяет сортировку по месяцам
func TestMonthSort(t *testing.T) {
	input := []string{
		"Mar",
		"Jan",
		"Feb",
	}

	expected := []string{
		"Jan",
		"Feb",
		"Mar",
	}

	// Включаем флаг -M
	monthFlag = new(bool)
	*monthFlag = true

	// Выполняем сортировку
	sort.SliceStable(input, func(i, j int) bool {
		return compareLines(input[i], input[j])
	})

	// Проверяем результат
	for i, line := range expected {
		if input[i] != line {
			t.Errorf("Ошибка на строке %d: ожидалось '%s', получено '%s'", i, line, input[i])
		}
	}
}
