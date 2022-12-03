package main

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
