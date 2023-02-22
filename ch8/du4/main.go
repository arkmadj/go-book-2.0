package main

import "os"

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select{
		case sema <-struct{}{}
		case <-done
	}
}