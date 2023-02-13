package main

import "net"

type conn struct {
	rw           net.Conn
	dataHostPort string
	prevCmd      string
	pasvListener net.Listener
	cmdErr       error
	binary       bool
}
