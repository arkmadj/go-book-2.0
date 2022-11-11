package main

import "fmt"

func removeAdjacents(strings []string) []string {
	w := 0
	for _, s := range strings {
		if strings[w] == s {
			continue
		}
		w++
		strings[w] = s
	}
	return strings[:w+1]
}

func main() {
	a := []string{"a", "b", "c", "c", "d", "a"}
	fmt.Println(removeAdjacents(a))
}
