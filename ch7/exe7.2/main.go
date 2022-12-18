package main

import "io"

type byteCounter struct {
	w       io.Writer
	written int64
}
