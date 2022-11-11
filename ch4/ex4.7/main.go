package main

import "fmt"

func reverse(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

func main() {
	a := []byte("Ahmad Jinadu")
	reverse(a)
	fmt.Println(string(a))
}
