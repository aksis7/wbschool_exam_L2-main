package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Определение флагов
	after := flag.Int("A", 0, "Print N lines after a match")
	before := flag.Int("B", 0, "Print N lines before a match")
	context := flag.Int("C", 0, "Print N lines before and after a match (overrides -A and -B)")
	count := flag.Bool("c", false, "Print only a count of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case distinctions")
	invert := flag.Bool("v", false, "Invert match (exclude matching lines)")
	fixed := flag.Bool("F", false, "Interpret pattern as a fixed string")
	lineNum := flag.Bool("n", false, "Print line numbers")

	flag.Parse()

	// Проверка наличия паттерна и файла
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run task.go [OPTIONS] PATTERN [FILE]")
		return
	}

	pattern := flag.Arg(0)
	var filename string
	if len(flag.Args()) > 1 {
		filename = flag.Arg(1)
	}

	// Чтение данных из файла или стандартного ввода
	var scanner *bufio.Scanner
	if filename != "" {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			return
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	// Подготовка паттерна
	var re *regexp.Regexp
	var err error
	if *fixed {
		pattern = regexp.QuoteMeta(pattern)
	}
	if *ignoreCase {
		pattern = "(?i)" + pattern
	}
	re, err = regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid pattern: %v\n", err)
		return
	}

	// Отладочные сообщения
	fmt.Println("DEBUG: Pattern:", pattern)

	// Чтение строк
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading lines: %v\n", err)
		return
	}

	// Отладочные сообщения
	fmt.Println("DEBUG: File contents:")
	for i, line := range lines {
		fmt.Printf("DEBUG: Line %d: %s\n", i+1, line)
	}

	// Поиск совпадений
	matchIndices := []int{}
	for i, line := range lines {
		matched := re.MatchString(line)
		if *invert {
			matched = !matched
		}
		if matched {
			matchIndices = append(matchIndices, i)
		}
	}

	// Отладочные сообщения
	fmt.Println("DEBUG: Match indices:", matchIndices)

	// Если включен флаг `-c`, выводим только количество совпадений
	if *count {
		fmt.Println(len(matchIndices))
		return
	}

	// Определение контекста
	beforeCtx := *before
	afterCtx := *after
	if *context > 0 {
		beforeCtx = *context
		afterCtx = *context
	}

	// Печать результатов с учетом контекста
	printedLines := make(map[int]bool)
	for _, idx := range matchIndices {
		start := idx - beforeCtx
		if start < 0 {
			start = 0
		}
		end := idx + afterCtx
		if end >= len(lines) {
			end = len(lines) - 1
		}
		for i := start; i <= end; i++ {
			if _, printed := printedLines[i]; !printed {
				if *lineNum {
					fmt.Printf("%d:%s\n", i+1, lines[i])
				} else {
					fmt.Println(lines[i])
				}
				printedLines[i] = true
			}
		}
	}
}
