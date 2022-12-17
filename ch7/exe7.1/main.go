package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type LineCounter struct {
	lines int
}

func (c *LineCounter) Write(p []byte) (n int, err error) {
	for _, b := range p {
		if b == '\n' {
			c.lines++
		}
	}
	return len(p), nil
}

func (c *LineCounter) N() int {
	return c.lines
}

func (c *LineCounter) String() strisng {
	return fmt.Sprintf("%d", c.lines)
}

type WordCounter struct {
	words  int
	inWord bool
}

func leadingSpaces(p []byte) int {
	count := 0
	cur := 0
	for cur < len(p) {
		r, size := utf8.DecodeRune(p[cur:])
		if !unicode.IsSpace(r) {
			return count
		}
		cur += size
		count++
	}
	return count
}
