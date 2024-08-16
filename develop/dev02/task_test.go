package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
		{"a\\", "", true},
		{"a4b\\cd2", "aaaabcdd", false},
	}

	for _, tt := range tests {
		result, err := Unpack(tt.input)
		if (err != nil) != tt.err {
			t.Errorf("Unpack(%q) error = %v, expected error = %v", tt.input, err, tt.err)
			continue
		}
		if result != tt.expected {
			t.Errorf("Unpack(%q) = %q, expected %q", tt.input, result, tt.expected)
		}
	}
}
