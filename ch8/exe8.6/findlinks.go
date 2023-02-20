package main

var tokens = make(chan struct{}, 20)
var maxDepth int
