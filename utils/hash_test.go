package utils

import (
	"testing"
)

func TestMD5Hex(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"12345", "827ccb0eea8a706c4c34a16891f84e7b"},
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
		{"world", "7d793037a0760186574b0282f2f435e7"},
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
	}
	for _, c := range cases {
		got := MD5Hex(c.input)
		if got != c.expected {
			t.Errorf("MD5Hex(%q) == %q, want %q", c.input, got, c.expected)
		}
	}
}
