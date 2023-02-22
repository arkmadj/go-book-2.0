package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type SizeResponse struct {
	root int
	size int64
}

var sema = make(chan struct{}, 20)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

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

func walkDir(dir string, n *sync.WaitGroup, root int, sizeResponses chan<- SizeResponse) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, root, sizeResponses)
		} else {
			sizeResponses <- SizeResponse{root, entry.Size()}
		}
	}
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, r := range roots {
		fmt.Printf("%10d files %.3f GB under %s\n", nfiles[i], float64(nbytes[i])/1e9, r)
	}
}
