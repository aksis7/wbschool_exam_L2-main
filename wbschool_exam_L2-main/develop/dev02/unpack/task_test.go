package unpack

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"qwe\\\\5", "qwe\\\\\\\\\\", false},
		{"a2\\", "", true}, // Некорректная строка
		{"\\1\\2\\3", "123", false},
		{"a\\5b", "a5b", false},
	}

	for _, tt := range tests {
		result, err := Unpack(tt.input)
		if (err != nil) != tt.hasError {
			t.Errorf("Unpack(%q) unexpected error: %v", tt.input, err)
		}
		if result != tt.expected {
			t.Errorf("Unpack(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
