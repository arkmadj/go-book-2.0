package main

import (
	"fmt"
)

func anagram(a, b string) bool {
	aFreq := make(map[rune]int)
	for _, c := range a {
		aFreq[c]++
	}
	bFreq := make(map[rune]int)
	for _, c := range b {
		bFreq[c]++
	}
	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}
	for k, v := range bFreq {
		if aFreq[k] != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%s - %s: %t\n", "ahmad", "Ahmadi", anagram("ahmad", "hamad"))
	fmt.Printf("%s - %s: %t\n", "Ahmad", "Hamad", anagram("ahmad", "hamad"))
	fmt.Printf("%s - %s: %t\n", "Ahmad", "Adekunle", anagram("ahmad", "hamad"))
}
