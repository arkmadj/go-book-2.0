package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/big"
	"math/cmplx"
	"net/http"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		epsX                   = (xmax - ymin) / width
		epsY                   = (ymax - ymin) / height
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		offX := []float64{-epsX, epsX}
		offY := []float64{-epsY, epsY}

		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				subPixels := make([]color.Color, 0)
				for i := 0; i < 2; i++ {
					for j := 0; j < 2; j++ {
						z := complex(x+offX[i], y+offY[j])
						subPixels = append(subPixels, mandelbrotBigFloat(z))
					}
				}
				img.Set(px, py, avg(subPixels))
			}
		}
		png.Encode(w, img)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func avg(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := len(colors)
	for _, c := range colors {
		r_, g_, b_, a_ := c.RGBA()
		r += uint16(r_ / uint32(n))
		g += uint16(g_ / uint32(n))
		b += uint16(b_ / uint32(n))
		a += uint16(a_ / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const constrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50:
				return color.RGBA{100, 0, 0, 255}
			default:
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}

func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	const constrast = 15
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			switch {
			case n > 50:
				return color.RGBA{100, 0, 0, 255}
			default:
				logScale := math.Log(float64(n) / math.Log(float64(iterations)))
				return color.RGBA{10, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200
	const constrast = 15

	zR := (&big.Float{}).SetFloat64(real(z))
	zI := (&big.Float{}).SetFloat64(imag(z))
	var vR, vI = &big.Float{}, &big.Float{}
	for i := uint8(0); i < iterations; i++ {
		vR2, vI2 := &big.Float{}, &big.Float{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Float{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Float{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Float{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewFloat(4)) == 1 {
			switch {
			case i > 50:
				return color.RGBA{100, 0, 0, 255}
			default:
				logScale := math.Log(float64(i)) / math.Log(float64(iterations))
				return color.RGBA{255, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}
