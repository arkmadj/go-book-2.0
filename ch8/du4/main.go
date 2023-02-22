package main

import "os"


var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select{
		case sema <-struct{}{}
		case <-done
	}
}