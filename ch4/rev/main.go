package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := []string{"0", "1", "2", "3", "4", "5"}
	c := []string{"0", "1", "2", "6", "4", "5"}
	d := []string{"0", "1", "2", "7", "4", "5"}
	reverse(a[:2])
	reverse(a[2:])
	reverse(a)
	fmt.Println(a)
	fmt.Println(equal(b, c))
	fmt.Println(equal(b, d))
	fmt.Println(equal(c, d))
}
