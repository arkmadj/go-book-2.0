package main

import (
	"image/color"
	"math"
	"math/cmplx"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func newtow(z complex128) color.Color {
	iterations := 37
	for n := uint8(0); int(n) < iterations; n++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(cmplx.Pow(z, 4)-1) < 1e-6 {
			return color.Gray{255 - uint8(math.Log(float64(n))/math.Log(float64(iterations+0))*244)}
		}
	}
	return color.Black
}
