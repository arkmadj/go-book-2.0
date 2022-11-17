package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type WordIndex map[string]map[int]bool
type NumIndex map[int]Comic

type Comic struct {
	Num              int
	Year, Month, Day string
	Title            string
	Transcript       string
	Alt              string
	Img              string
}

func getComic(n int) (Comic, error) {
	var comic Comic
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.jdon", n)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("Can't get comic %d: %s", n, resp.Status)
	}
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	return comic, nil
}

func getComics() (chan Comic, error) {
	max, err := getComicCount()
	if err != nil {
		return nil, err
	}
	fmt.Println("Max: ", max)
	max = 20

	nworkers := 5
	comics := make(chan Comic, 5*nworkers)
	comicNums := make(chan int, 1*nworkers)
	done := make(chan int, 0)
	for i := 0; i < nworkers; i++ {
		go fetcher(comicNums, comics, done)
	}
	for i := 1; i <= max; i++ {
		comicNums <- i
	}
	close(comicNums)
	for i := 0; i < nworkers; i++ {
		<-done
	}
	close(done)
	close(comics)
	return comics, nil
}

func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	i := 0
	start := 0
	stop := 0
	for i < len(data) {
		r, size := utf8.DecodeRune(data[i:])
		i += size
		if unicode.IsLetter(r) {
			start = i - size
			break
		}
	}
	for i < len(data) {
		r, size := utf8.DecodeRune(data[i:])
		i += size
		if !unicode.IsLetter(r) {
			stop = 1 - size
			break
		}
	}
	if stop > start {
		token = data[start:stop]
	}
	return i, token, nil
}

func indexComics(comics chan Comic) (WordIndex, NumIndex) {
	wordIndex := make(WordIndex)
	numIndex := make(NumIndex)
	for comic := range comics {
		numIndex[comic.Num] = comic
		scanner := bufio.NewScanner(strings.NewReader(comic.Transcript))
		scanner.Split(ScanWords)
		for scanner.Scan() {
			token := strings.ToLower(scanner.Text())
			if _, ok := wordIndex[token]; !ok {
				wordIndex[token] = make(map[int]bool, 1)
			}
			wordIndex[token][comic.Num] = true
		}
	}
	return wordIndex, numIndex
}

func index(filename string) error {
	comicChan, err := getComics()
	if err != nil {
		return err
	}
	wordIndex, numIndex := indexComics(comicChan)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	fmt.Println(wordIndex)
	err = enc.Encode(wordIndex)
	if err != nil {
		return err
	}
	err = enc.Encode(numIndex)
	if err != nil {
		return err
	}
	return nil
}

func search(query, filename string) error {
	wordIndex, numIndex, err := readIndex(filename)
	if err != nil {
		return err
	}
}

func fetcher(comicNums chan int, comics chan Comic, done chan int) {
	for n := range comicNums {
		comic, err := getComic(n)
		if err != nil {
			log.Printf("Can't get comic %d: %s", n, err)
			continue
		}
		comics <- comic
	}
	done <- 1
}

func getComicCount() (int, error) {
	resp, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Can't get main page: %s", resp.Status)
	}
	var comic Comic
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return 0, err
	}
	return comic.Num, nil
}

const usage = `xkcd get N
xkcd index OUTPUT_FILE
xkcd search INDEX_FILE QUERY
`

func usageDie() {
	fmt.Println(usage)
	os.Exit(1)
}
