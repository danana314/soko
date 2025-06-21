package util

import (
	"regexp"
	"testing"
)

func TestPrintStruct(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name:     "string",
			input:    "hello",
			expected: `"hello"`,
		},
		{
			name:     "int",
			input:    42,
			expected: "42",
		},
		{
			name:     "struct",
			input:    struct{ Name string }{Name: "test"},
			expected: `struct { Name string }{Name:"test"}`,
		},
		{
			name:     "slice",
			input:    []int{1, 2, 3},
			expected: "[]int{1, 2, 3}",
		},
		{
			name:     "map",
			input:    map[string]int{"a": 1, "b": 2},
			expected: `map[string]int{"a":1, "b":2}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PrintStruct(tt.input)
			if result != tt.expected {
				t.Errorf("PrintStruct(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGenerateID(t *testing.T) {
	t.Run("generates non-empty ID", func(t *testing.T) {
		id := GenerateID()
		if id == "" {
			t.Error("GenerateID() returned empty string")
		}
	})

	t.Run("generates unique IDs", func(t *testing.T) {
		ids := make(map[string]bool)
		for range 100 {
			id := GenerateID()
			if ids[id] {
				t.Errorf("GenerateID() generated duplicate ID: %s", id)
			}
			ids[id] = true
		}
	})

	t.Run("generates valid shortuuid format", func(t *testing.T) {
		id := GenerateID()
		// shortuuid should be alphanumeric and around 22 characters
		matched, err := regexp.MatchString(`^[A-Za-z0-9]{20,25}$`, id)
		if err != nil {
			t.Errorf("regex error: %v", err)
		}
		if !matched {
			t.Errorf("GenerateID() = %q, doesn't match expected shortuuid format", id)
		}
	})
}