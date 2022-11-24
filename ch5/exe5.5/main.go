package main

import (
	"fmt"
	"net/http"

	"github.com/ahmad/go-book-2.0/ch4/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("Parsing HTML: %s", err)
		return
	}
	words, images = CountWordsAndImages(doc)
	return
}
