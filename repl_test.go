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
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    JELLO is great _ t       ",
			expected: []string{"jello", "is", "great", "_", "t"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length does not match")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words do not match")
				break

			}

		}
	}
}
