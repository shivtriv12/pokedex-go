package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  Hello  World  ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedword := c.expected[i]
			if word != expectedword {
				t.Errorf("%v is expected but got %v.Test Failed", expectedword, word)
			}
		}
	}
}
