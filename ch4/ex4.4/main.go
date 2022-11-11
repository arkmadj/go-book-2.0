package main

import "fmt"

func rotate_ints(ints []int, positions int) {
	first := [3]int{}
	for a, i := range ints[:positions] {
		first[a] = i
	}
	copy(ints, ints[positions:])
	for i := 1; i <= positions; i++ {
		ints[len(ints)-i] = first[positions-i]
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate_ints(s, 3)
	fmt.Println(s)
}
