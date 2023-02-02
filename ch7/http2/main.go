package main

import "fmt"

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}
