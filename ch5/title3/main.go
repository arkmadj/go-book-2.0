package main

import "golang.org/x/net/html"

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
}
