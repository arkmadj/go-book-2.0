package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/metal3d/go-slugify"
)

const APIURL = "http://www.omdbapi.com/?"

type Movie struct {
	Title  string
	Year   string
	Poster string
}

func (m Movie) posterFilename() string {
	ext := filepath.Ext(m.Poster)
	title := slugify.Marshal(m.Title)
	return fmt.Sprintf("%s_(%s)%s", title, m.Year, ext)
}

func getMovie(title string) (movie Movie, err error) {
	url_ := fmt.Sprintf("%st=%s", APIURL, url.QueryEscape(title))
	resp, err := http.Get(url_)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("%d response from %s", resp.StatusCode, url_)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		return
	}
	return
}
