package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file := os.Args[1]
	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}
	wordFreq := make(map[string]int)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		wordFreq[word]++
	}
	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(1)
	}
	for word, n := range wordFreq {
		fmt.Printf("%-30s %d\n", word, n)
	}
}
