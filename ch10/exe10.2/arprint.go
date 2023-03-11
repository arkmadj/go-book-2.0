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
