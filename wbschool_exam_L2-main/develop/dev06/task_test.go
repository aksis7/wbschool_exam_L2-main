package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name       string
		input      string // Входные данные
		delimiter  string // Разделитель
		fields     string // Поля
		separated  bool   // Флаг separated
		expected   string // Ожидаемый результат
	}{
		{
			name:      "Basic Tab-separated columns",
			input:     "a\tb\tc\nd\te\tf\n",
			delimiter: "\t",
			fields:    "1,3",
			separated: false,
			expected:  "a\tc\nd\tf\n",
		},
		{
			name:      "Custom Delimiter Colon",
			input:     "a:b:c\nd:e:f\n",
			delimiter: ":",
			fields:    "1,2",
			separated: false,
			expected:  "a:b\nd:e\n",
		},
		{
			name:      "Separated flag with missing delimiter",
			input:     "a:b:c\nno delimiter here\n",
			delimiter: ":",
			fields:    "1,2",
			separated: true,
			expected:  "a:b\n",
		},
		{
			name:      "Invalid Field Index",
			input:     "a\tb\tc\n",
			delimiter: "\t",
			fields:    "5",
			separated: false,
			expected:  "\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := strings.NewReader(test.input)
			var output bytes.Buffer

			runCut(input, &output, test.fields, test.delimiter, test.separated)

			got := output.String()
			if got != test.expected {
				t.Errorf("\nExpected:\n%q\nGot:\n%q", test.expected, got)
			}
		})
	}
}
