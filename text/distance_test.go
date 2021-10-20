package text

import (
	"testing"
)

func TestLevensthein(t *testing.T) {
	var tests = []struct {
		s1       string
		s2       string
		expected int
	}{
		{"kitten", "sitting", 3},
		{"benyam", "ephrem", 5},
		{"a", "a", 0},
		{"a", "b", 1},
	}

	for _, test := range tests {
		if output := Levensthein(test.s1, test.s2); output != test.expected {
			t.Error("Test Failed,", test.expected, " expected,", output, " received.")
		}
	}
}
