package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

var ErrNotExist = errors.New("file does not exist")

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func IsNotExist(err error) bool {
	if pe, ok := err.(*PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err))
}
