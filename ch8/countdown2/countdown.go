package main

import (
	"fmt"
	"os"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
}

func launch() {
	fmt.Println("Lift off!")
}
