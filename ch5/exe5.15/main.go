package main

import "fmt"

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

func max2(first int, nums ...int) int {
	m := first
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(min(values...))
	fmt.Println(max(values...))
	fmt.Println(min2(values...))
	fmt.Println(max2(6, values...))
}
