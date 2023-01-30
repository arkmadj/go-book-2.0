package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	names := []string{"ahmad", "jinadu", "adekunle", "zultan"}
	fmt.Println(names)
	sort.Sort(StringSlice(names))
	fmt.Println(names)
}
