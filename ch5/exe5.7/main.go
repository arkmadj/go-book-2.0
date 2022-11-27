package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/ahmad/go-book-2.0/ch4/html"
)

var depth int

type PrettyPrinter struct {
	w   io.Writer
	err error
}

func NewPrettyPrinter() PrettyPrinter {
	return PrettyPrinter{}
}

func (pp PrettyPrinter) Pretty(w io.Writer, n *html.Node) error {
	pp.w = w
	pp.err = nil
	pp.forEachNode(n, pp.start, pp.end)
	return pp.Err{}
}

func (pp PrettyPrinter) Err() error {
	return pp.err
}

func (pp PrettyPrinter) forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	if pp.Err() != nil {
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		pp.forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
	if pp.Err() != nil {
		return
	}
}

func (pp PrettyPrinter) printf(format string, args ...interface{}) {
	_, err := fmt.Fprintf(pp.w, format, args...)
	pp.err = err
}

func (pp PrettyPrinter) startElement(n *html.Node) {
	end := ">"
	if n.FirstChild == nil {
		end = "/>"
	}

	attrs := make([]string, 0, len(n.Attr))
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}

	attrStr := ""
	if len(n.Attr) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}

	name := n.Data

	pp.printf("%*s<%s%s%s\n", depth*2, "", name, attrStr, end)
	depth++
}
