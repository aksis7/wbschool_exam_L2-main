package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Флаги командной строки
var (
	columnFlag   = flag.Int("k", 0, "Указать колонку для сортировки (начиная с 1)")
	numericFlag  = flag.Bool("n", false, "Сортировать по числовому значению")
	reverseFlag  = flag.Bool("r", false, "Сортировать в обратном порядке")
	uniqueFlag   = flag.Bool("u", false, "Не выводить повторяющиеся строки")
	monthFlag    = flag.Bool("M", false, "Сортировать по названию месяца")
	ignoreSpace  = flag.Bool("b", false, "Игнорировать хвостовые пробелы")
	checkSorted  = flag.Bool("c", false, "Проверить, отсортированы ли данные")
	humanNumeric = flag.Bool("h", false, "Сортировать по числовым значениям с суффиксами (K, M, G, T)")
)

func getMonthOrder(month string) int {
	months := map[string]int{
		"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4,
		"May": 5, "Jun": 6, "Jul": 7, "Aug": 8,
		"Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
	}

	if order, ok := months[month]; ok {
		return order
	}
	return -1 // Возвращаем -1, если месяц не распознан
}

// Сопоставление месяцев для флага -M

// parseHumanReadableNumber преобразует числа с суффиксами (K, M, G, T) в числовое значение

func parseHumanReadableNumber(s string) float64 {
	suffixes := map[byte]float64{
		'K': 1e3,
		'M': 1e6,
		'G': 1e9,
		'T': 1e12,
	}

	if len(s) == 0 {
		return 0
	}

	lastChar := s[len(s)-1]
	if mul, ok := suffixes[lastChar]; ok {
		num, err := strconv.ParseFloat(s[:len(s)-1], 64)
		if err == nil {
			return num * mul
		}
	}

	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return math.NaN()
	}
	return num
}

func compareLines(a, b string) bool {
	colsA := strings.Fields(a)
	colsB := strings.Fields(b)

	col := *columnFlag - 1
	if col < 0 {
		col = 0
	}

	fieldA, fieldB := "", ""
	if col < len(colsA) {
		fieldA = colsA[col]
	}
	if col < len(colsB) {
		fieldB = colsB[col]
	}

	// Сортировка по месяцам (-M)
	if *monthFlag {
		monthA := getMonthOrder(fieldA)
		monthB := getMonthOrder(fieldB)

		if monthA != -1 && monthB != -1 {
			return monthA < monthB
		}
	}

	// Числовая сортировка с учетом суффиксов (-h)
	if *humanNumeric {
		numA := parseHumanReadableNumber(fieldA)
		numB := parseHumanReadableNumber(fieldB)

		if math.IsNaN(numA) || math.IsNaN(numB) {
			return fieldA < fieldB
		}
		return numA < numB
	}

	// Числовая сортировка (-n)
	if *numericFlag {
		numA, errA := strconv.ParseFloat(fieldA, 64)
		numB, errB := strconv.ParseFloat(fieldB, 64)
		if errA == nil && errB == nil {
			return numA < numB
		}
	}

	// Лексикографическая сортировка по умолчанию
	return fieldA < fieldB
}

// checkIfSorted проверяет, отсортирован ли массив
func checkIfSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if !compareLines(lines[i-1], lines[i]) {
			return false
		}
	}
	return true
}

// removeDuplicates удаляет дубликаты из слайса строк
func removeDuplicates(lines []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			result = append(result, line)
		}
	}
	return result
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: go run main.go [OPTIONS] <input_file>")
		return
	}

	filePath := flag.Arg(0)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if *ignoreSpace {
			line = strings.TrimRight(line, " ")
		}
		lines = append(lines, line)
	}

	if *checkSorted {
		if checkIfSorted(lines) {
			fmt.Println("Файл отсортирован.")
		} else {
			fmt.Println("Файл не отсортирован.")
		}
		return
	}

	// Удаляем дубликаты
	if *uniqueFlag {
		lines = removeDuplicates(lines)
	}

	// Сортировка строк
	sort.SliceStable(lines, func(i, j int) bool {
		return compareLines(lines[i], lines[j])
	})

	// Реверс сортировки
	if *reverseFlag {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}

	// Запись в файл
	outputFile, err := os.Create("sorted_output.txt")
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}
	defer outputFile.Close()

	for _, line := range lines {
		fmt.Fprintln(outputFile, line)
	}

	fmt.Println("Сортировка завершена. Результат сохранен в sorted_output.txt")
}
