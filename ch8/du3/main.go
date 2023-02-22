package main

var sema = make(chan chan struct{}, 20)
