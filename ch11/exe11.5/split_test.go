package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s, sep string
		want   int
	}{
		{"a:b:c", ":", 3},
		{"1,2,3,4,5", ",", 5},
		// {"x y z", " ", 4},
	}
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, test.want)
		}
	}
}
