package main

import "math"

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
)

var sin30, cos30 = 0.5, math.Sqrt(3.0 / 4.0)
