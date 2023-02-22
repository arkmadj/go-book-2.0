package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type SizeResponse struct {
	root int
	size int64
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() {
		<-sema
	}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func walkDir(dir string, n *sync.WaitGroup, root int, sizeResponses chan<- SizeResponse)
