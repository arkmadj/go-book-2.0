package main

type client struct {
	Out  chan<- string
	Name string
}
