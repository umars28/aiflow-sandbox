package greet

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		salutation string
		want       string
	}{
		{"populated name", "Ada", "Hello", "Hello, Ada"},
		{"empty name", "", "Hello", "Hello"},
		{"whitespace-only name", "   \t ", "Hello", "Hello"},
		{"custom salutation", "Ada", "Howdy", "Howdy, Ada"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.input, tt.salutation); got != tt.want {
				t.Errorf("Greet(%q, %q) = %q, want %q", tt.input, tt.salutation, got, tt.want)
			}
		})
	}
}
