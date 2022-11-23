package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ahmad/go-book-2.0/ch4/html"
)

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
