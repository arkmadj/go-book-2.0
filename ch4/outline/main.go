package main

import (
	"fmt"
	"os"

	"github.com/ahmad/go-book-2.0/ch4/html"
)

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func main(){
	doc, err := html.Parse(os.Stdin)
	if err !- nil {
		fmt.Fprintf(os.Stderr, "Outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
