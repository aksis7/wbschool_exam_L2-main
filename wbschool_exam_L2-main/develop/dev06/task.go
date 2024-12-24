package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Определение флагов
	fields := flag.String("f", "", "Выбрать поля (колонки), например: 1,2,3")
	delimiter := flag.String("d", "\t", "Использовать другой разделитель (по умолчанию TAB)")
	separated := flag.Bool("s", false, "Только строки с разделителем")

	flag.Parse()

	if *fields == "" {
		fmt.Fprintln(os.Stderr, "Ошибка: флаг -f обязателен")
		flag.Usage()
		os.Exit(1)
	}

	// Запуск основной логики
	runCut(os.Stdin, os.Stdout, *fields, *delimiter, *separated)
}

// runCut выполняет основную логику утилиты cut
func runCut(input io.Reader, output io.Writer, fields string, delimiter string, separated bool) {
	fieldIndexes := parseFields(fields)

	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()

		// Проверка наличия разделителя
		if separated && !strings.Contains(line, delimiter) {
			continue
		}

		// Разделяем строку по разделителю
		columns := strings.Split(line, delimiter)

		// Выводим запрашиваемые колонки
		result := extractFields(columns, fieldIndexes, delimiter)

		// Если результат пустой, явно добавляем новую строку
		if result == "" {
			writer.WriteString("\n")
		} else {
			writer.WriteString(result + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения ввода: %v\n", err)
	}
}

// parseFields разбирает строку с номерами полей
func parseFields(fields string) []int {
	parts := strings.Split(fields, ",")
	var indexes []int
	for _, part := range parts {
		var index int
		_, err := fmt.Sscanf(part, "%d", &index)
		if err != nil || index < 1 {
			fmt.Fprintf(os.Stderr, "Некорректный номер поля: %s\n", part)
			os.Exit(1)
		}
		indexes = append(indexes, index-1)
	}
	return indexes
}

// extractFields выбирает указанные поля с разделителем
func extractFields(columns []string, indexes []int, delimiter string) string {
	var result []string
	for _, index := range indexes {
		if index < len(columns) {
			result = append(result, columns[index])
		}
	}
	return strings.Join(result, delimiter)
}
