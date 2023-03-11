package zip

import (
	"archive/zip"
	"io"
)

type reader struct {
	zipReader *zip.Reader
	filesLeft []*zip.File
	r         io.ReadCloser
	toWrite   string
}
