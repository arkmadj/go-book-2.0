package main

import "fmt"

type LineCounter struct {
	lines int
}

func (c *LineCounter) Write(p []byte) (n int, err error) {
	for _, b := range p {
		if b == '\n' {
			c.lines++
		}
	}
	return len(p), nil
}

func (c *LineCounter) N() int {
	return c.lines
}

func (c *LineCounter) String() string {
	return fmt.Sprintf("%d", c.lines)
}
