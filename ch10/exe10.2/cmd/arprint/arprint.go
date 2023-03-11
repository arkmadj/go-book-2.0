package main

import (
	"fmt"
	"io"
	"os"

	arprint "github.com/ahmad/go-book-2.0/ch10/exe10.2"
)

func printArchive(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := arprint.Open(f)
	if err != nil {
		return fmt.Errorf("open archive reader: %s", err)
	}
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		return fmt.Errorf("printing: %s", err)
	}
	return nil
}
