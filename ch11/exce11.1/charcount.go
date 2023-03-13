package main

import (
	"bufio"
	"io"
	"log"
	"unicode"
)

func charCount(r io.Reader) (runes map[rune]int, props map[string]int, sizes map[int]int, invalid int) {
	runes = make(map[rune]int)
	props = make(map[string]int)
	sizes = make(map[int]int)
	invalid = 0

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("charCount: %s", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for name, rangeTable := range unicode.Properties {
			if unicode.In(r, rangeTable) {
				props[name]++
			}
		}
		runes[r]++
		sizes[n]++
	}
}
