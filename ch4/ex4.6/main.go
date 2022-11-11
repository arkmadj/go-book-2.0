package main

import (
	"unicode"
	"unicode/utf8"
)

func removeAdjacentsSpaces(b []byte) {
	w := 0
	r, _ := utf8.DecodeRune(b)
	for _, s := range r {
		if unicode.IsSpace(r) {
			continue
		}
		w++
		b[w] = s
	}
	return b[:w+1]
}
