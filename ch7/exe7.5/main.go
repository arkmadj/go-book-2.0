package main

import "io"

type limitReader struct {
	r        io.Reader
	n, limit int
}
