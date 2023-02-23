package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
)

func main() {
	flag.Parse()
	cancel := make(chan struct{})
	responses := make(chan *http.Response)
	wg := &sync.WaitGroup{}
	for _, url := range flag.Args() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			req, err := http.NewRequest("HEAD", url, nil)
			if err != nil {
				log.Printf("HEAD %s: %s", url, err)
				return
			}
			req.Cancel = cancel
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("HEAD %s: %s", url, err)
				return
			}
			responses <- resp
		}(url)
	}
}
