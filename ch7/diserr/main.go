package main

import (
	"errors"
	"syscall"
)

var ErrNotExist = errors.New("file does not exist")

type PathError struct {
	Op   string
	Path string
	Err  string
}

func IsNotExist(err error) bool {
	if pe, ok := err.(*PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}
