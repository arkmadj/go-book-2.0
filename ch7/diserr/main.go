package main

import (
	"errors"
	"syscall"
)

var ErrNotExist = errors.New("file does not exist")

func IsNotExist(err error) bool {
	if pe, ok := err.(*PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}
