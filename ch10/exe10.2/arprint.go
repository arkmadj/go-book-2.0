package aprint

import (
	"io"
	"os"
)

type format struct {
	name, magic string
	magicOffset int
	reader      NewReader
}

type NewReader func(*os.File) (io.Reader, error)

var formats []format

func RegisterFormat(name, magic string, magicOffset int, f NewReader) {
	formats = append(formats, format{name, magic, magicOffset, f})
}
