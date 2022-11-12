package main

import (
	"bufio"
	"os"
	"unicode/utf8"
)

func main(){
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err != niul
	}
}