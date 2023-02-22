package main

import "os"

var done = make(chan struct{})

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select{
		case sema <-struct{}{}
		case <-done
	}
}