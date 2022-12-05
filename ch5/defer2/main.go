package main

import (
	"os"
	"runtime"
)

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
