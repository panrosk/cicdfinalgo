package tests

import (
	"cicdfinalgo/utils"
	"testing"
)

func TestSayHiTo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Greet World", "World", "Hello, World!"},
		{"Greet User", "Oscar", "Hello, Oscar!"},
		{"Greet Empty", "", "Hello, !"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.SayHiTo(tt.input)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
