package main

import (
	"fmt"
	"sort"
	"strings"
)

// Функция для сортировки букв в слове
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

// Функция поиска множеств анаграмм
func FindAnagramSets(words []string) map[string][]string {
	anagramGroups := make(map[string][]string)
	uniqueWords := make(map[string]bool)

	// Приведение слов к нижнему регистру и фильтрация уникальных слов
	for _, word := range words {
		lowered := strings.ToLower(word)
		uniqueWords[lowered] = true
	}

	// Группировка анаграмм
	for word := range uniqueWords {
		sortedWord := sortString(word)
		anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
	}

	// Формирование результата
	result := make(map[string][]string)
	for _, group := range anagramGroups {
		if len(group) > 1 {
			sort.Strings(group)
			result[group[0]] = group
		}
	}

	return result
}

// Пример использования
func main() {
	words := []string{
		"пятак", "пятка",
		"листок", "слиток", "столик",
		"яблоко", "кот", "ток", "окт", "тяпка",
	}

	anagramSets := FindAnagramSets(words)

	for key, group := range anagramSets {
		fmt.Printf("%s: %v\n", key, group)
	}
}
