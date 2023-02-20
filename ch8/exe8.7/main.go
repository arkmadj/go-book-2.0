package main

import (
	"net/url"
	"sync"
)

var tokens = make(chan struct{}, 20)
var maxDepth int
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}
var base *url.URL

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	tokens <- struct{}{}
	urls, err := visit(url)
}
