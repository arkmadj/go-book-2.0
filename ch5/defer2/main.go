package main

import (
	"fmt"
	"os"
	"runtime"
)

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("deffer %d\n", x)
	f(x - 1)
}

func main() {
	defer printStack()
	f(3)
}
