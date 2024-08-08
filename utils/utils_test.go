package utils

import (
	"testing"
)

func TestIsPassCorrect(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{

		{"12345678", false},
		{"@#$%^&*", false},
		{"pass1@22", true},
		{"pass1@", false},
		{"short1@", false},
		{"longenough1$", true},
	}

	for _, test := range tests {
		result := IsPassCorrect(test.password)
		if result != test.expected {
			t.Errorf("IsPassCorrect(%q) = %v; expected %v", test.password, result, test.expected)
		}
	}
}
