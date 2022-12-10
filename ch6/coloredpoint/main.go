package main

import (
	"fmt"
	"image/color"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	m = nil
	fmt.Println(m.Get("item"))
	m.Add("item", "3")

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{255, 0, 0, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{1, 1}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}
