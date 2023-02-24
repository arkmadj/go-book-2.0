package main

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	mesages  = make(chan string)
)
