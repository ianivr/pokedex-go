package main

import "testing"

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: "  hello  world  ",
			want:  []string{"hello", "world"},
		},
		{
			input: "   foo   bar   baz   ",
			want:  []string{"foo", "bar", "baz"},
		},
		{
			input: "   leading and trailing spaces   ",
			want:  []string{"leading", "and", "trailing", "spaces"},
		},
	}

	for _, c := range tests {
		actual := cleanInput(c.input)
		if len(actual) != len(c.want) {
			t.Errorf("expected length: %d, got: %d", len(c.want), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.want[i]
			if word != expectedWord {
				t.Errorf("expected: %s, got: %s", expectedWord, word)
			}
		}
	}
}
