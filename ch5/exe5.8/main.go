package main

import "golang.org/x/net/html"

func ElementByID(n *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}
		return true
	}
	return forEachElement(n, pre, nil)
}
