package main

import (
	"fmt"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"
const APIURL = "https://api.github.com"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"createdat"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func (i Issue) CacheURL() string {
	return fmt.Sprintf("/issues/%d", i.Number)
}
