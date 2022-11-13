package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	a := [...]string{"a", "b", "c", "d", "e", "f"}
	for i := 0; i < len(a)-1; i++ {
		addEdge(a[i], a[i+1])
	}
	for i := 0; i < len(a)-1; i++ {
		fmt.Println(hasEdge(a[i], a[i+1]))
	}
	for i := 0; i < len(a)-2; i++ {
		fmt.Println(hasEdge(a[i], a[i+2]))
	}
}
