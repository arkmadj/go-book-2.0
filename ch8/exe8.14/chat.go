package main

import "time"

const timeout = 10 * time.Second

type client struct {
	Out  <-chan string
	Name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)
