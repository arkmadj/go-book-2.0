package main

import (
	"io"

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

func (pp PrettyPrinter) for
