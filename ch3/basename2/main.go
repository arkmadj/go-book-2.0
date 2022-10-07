package main

import (
	"fmt"
	"os"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	for _, input := range os.Args[1:] {
		fmt.Println(basename(input))
	}
}
