package main

import "flag"

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", "", "separator")
)
