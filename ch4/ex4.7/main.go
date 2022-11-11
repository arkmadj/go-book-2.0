package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

func revUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeLastRune((b[1:]))
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
	return b
}

func main() {
	a := []byte("Ahmad Jinadu")
	b := []byte("Monkey D. Luffy")
	reverse(a)
	revUTF8(b)
	fmt.Println(string(a))
	fmt.Println(string(b))
}
