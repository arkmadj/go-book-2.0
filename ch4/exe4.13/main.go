package main

import "path/filepath"

const APIURL = "http://www.omdbapi.com/?"

type Movie struct {
	Title  string
	Year   string
	Poster string
}

func (m Movie) posterFilename() string {
	ext := filepath.Ext(m.Poster)
}
