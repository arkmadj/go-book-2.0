package main

import (
	"flag"
	"io"
	"os"
)

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", "", "separator")
)

var out io.Writer = os.Stdout
