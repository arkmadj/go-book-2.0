package main

import (
	"fmt"
	"log"

	"github.com/ahmad/go-book-2.0/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
