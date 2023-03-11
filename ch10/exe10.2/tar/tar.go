package tar

import (
	"archive/tar"
	"os"
)

type reader struct {
	tarReader *tar.Reader
	file      *os.File
	toWrite   string
}

func (r *reader) Read(b []byte) (written int, err error) {
	for len(b) > 0 {
		if len(r.toWrite) > 0 {
		}
	}
}
