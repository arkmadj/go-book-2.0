package main

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
