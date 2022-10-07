package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ahmad/go-book/ch2/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		}

		f := unitconv.Fahrenheit(t)
		c := unitconv.Celsius(t)
		m := unitconv.Meter(t)
		ft := unitconv.Foot(t)
		p := unitconv.Pound(t)
		k := unitconv.Kilogram(t)

		fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n%s - %s, %s = %s\n", c, unitconv.CToF(c), f, unitconv.FToC(f), m, unitconv.MToFt(m), ft, unitconv.FtToM(ft), p, unitconv.PToK(p), k, unitconv.KToP(k))
	}
}
