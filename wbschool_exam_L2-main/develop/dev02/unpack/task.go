package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// Unpack распаковывает строку согласно заданному формату.
func Unpack(input string) (string, error) {
	var result strings.Builder
	var prevRune rune
	escaped := false

	for _, r := range input {
		if escaped {
			result.WriteRune(r)
			escaped = false
			prevRune = r
			continue
		}

		if r == '\\' {
			escaped = true
			continue
		}

		if unicode.IsDigit(r) {
			if prevRune == 0 {
				return "", errors.New("invalid string: digit cannot appear at the start")
			}

			count, err := strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}

			result.WriteString(strings.Repeat(string(prevRune), count-1))
			prevRune = 0
			continue
		}

		result.WriteRune(r)
		prevRune = r
	}

	if escaped {
		return "", errors.New("invalid string: unfinished escape sequence")
	}

	return result.String(), nil
}
