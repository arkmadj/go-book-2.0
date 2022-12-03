package main

import "golang.org/x/net/html"

func ElementsByTag(n *html.Node, tags ...string) []*html.Node {
	nodes := make([]*html.Node, 0)
	keep := make(map[string]bool, len(tags))
	for _, t := range tags {
		keep[t] = true
	}

	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		_, ok := keep[n.Data]
		if ok {
			nodes = append(nodes, n)
		}
		return true
	}
	forEachElementNode(n, pre, nil)
	return nodes
}

func forEachElementNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	u := make([]*html.Node, 0)
	u = append(u, n)
	for len(u) > 0 {
		n = u[0]
		u = u[1:]
		if pre != nil {
			if !pre(n) {
				return n
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
		if post != nil {
			if !post(n) {
				return n
			}
		}
	}
	return nil
}
