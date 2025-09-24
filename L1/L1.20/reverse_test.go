package main

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"snow dog sun", "sun dog snow"},
		{"hello", "hello"},
		{"a b c d", "d c b a"},
		{"", ""},
		{"one two", "two one"},
		{" one with space", "space with one "},
		{" space two ", " two space "},
	}

	for _, test := range tests {
		r := []rune(test.input)
		reverseWords(r, ' ')
		got := string(r)

		if got != test.expected {
			t.Errorf("reverseWords(%q): got %q, expected %q\n", test.input, got, test.expected)
		}
	}
}

func TestReverseRunes(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"wildberries", "seirrebdliw"},
		{"abc", "cba"},
		{" instance", "ecnatsni "},
		{" spaces ", " secaps "},
		{"", ""},
		{"a", "a"},
		{"абв", "вба"},
	}

	for _, test := range tests {
		r := []rune(test.input)
		reverseRunes(r)
		got := string(r)

		if got != test.expected {
			t.Errorf("reverseRunes(%q): got %q, expected %q\n", test.input, got, test.expected)
		}
	}
}
