package main

import "io"

type stringReaader struct {
	s string
}

func (r *stringReaader) Read(p []byte) (n int, err error){
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0{
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &stringReaader{s}
}

func main(){}