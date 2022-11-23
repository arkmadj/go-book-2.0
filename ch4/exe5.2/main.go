package main

import (
	"io"
	"os"

	"github.com/ahmad/go-book-2.0/ch4/html"
)

func tagFreq(r io.Reader) (map[string]int, error) {
	freq := make(map[string]int, 0)
	z := html.NewTokenizer(os.Stdin)
	var err error
	for {
		type_ := z.Next()
		if type_ == html.ErrorToken {
			break
		}
		name, _ := z.TagName()
		if len(name) > 0 {
			freq[string(name)]++
		}
	}
	if err != io.EOF {
		return freq, err
	}
	return freq, nil
}
