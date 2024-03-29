package aprint_test

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	. "github.com/ahmad/go-book-2.0/ch10/exe10.2"
	_ "github.com/ahmad/go-book-2.0/ch10/exe10.2/tar"
	_ "github.com/ahmad/go-book-2.0/ch10/exe10.2/zip"
)

func TestOpen(t *testing.T) {
	for _, archive := range []string{"rah.zip", "rah.tar"} {
		b := &bytes.Buffer{}
		f, err := os.Open(filepath.Join("testData", archive))
		if err != nil {
			t.Error(archive, err)
		}
		r, err := Open(f)
		if err != nil {
			t.Error(archive, err)
		}
		_, err = io.Copy(b, r)
		if err != nil {
			t.Error(archive, err)
		}
		want := `rah/b:
		contentsB
		rah/a:
		contentsA
		`

		got := b.String()
		if got != want {
			t.Errorf("%s: got %q, want %q", archive, got, want)
		}
	}
}
