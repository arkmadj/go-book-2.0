package main

import "io"

var depth int

type PrettyPrinter struct {
	w   io.Writer
	err error
}
