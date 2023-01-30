package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{3., 1, 4, 1}
	fmt.Println(sort.IntsAreSorted((values)))
	sort.Ints(values)
	fmt.Println(values)
}
