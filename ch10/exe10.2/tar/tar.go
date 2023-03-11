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
