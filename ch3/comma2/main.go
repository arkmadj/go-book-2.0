package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	b := &bytes.Buffer{}
	pre := len(s) % 3
	if pre == 0 {
		pre = 3
	}
	b.WriteString(s[:pre])
	for i := pre; i < len(s); i += 3 {
		b.WriteString(",")
		b.WriteString(s[i : i+3])
	}
	return b.String()
}

func main() {
	for _, input := range os.Args[1:] {
		fmt.Printf("%s\n", comma(input))
	}
}
