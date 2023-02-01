package main

import (
	"fmt"
	"sort"
)

func equal(i, j int, s sort.Interface) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

func IsPalindrome(s sort.Interface) bool {
	max := s.Len() - 1
	for i := 0; i < s.Len()/2; i++ {
		if !equal(i, max-i, s) {
			return false
		}
	}
	return true
}

func main() {
	intsA := []int{1, 2, 3, 4, 3, 2, 1}
	fmt.Printf("%v is palindrome: %t\n", intsA, IsPalindrome(sort.IntSlice(intsA)))

	intsB := []int{1, 2, 3, 4, 3, 2, 1}
	fmt.Printf("%v is palindrome: %t\n", intsB, IsPalindrome(sort.IntSlice(intsB)))

}
