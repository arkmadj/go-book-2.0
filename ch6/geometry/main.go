package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

type Path []Point

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

// func (path Path) TranslateBy(offset Point, add bool){
// 	var op func(p, q Point)Point
// 	if add {
// 		op = Point.
// 	}
// }

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	var origin Point

	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))
	fmt.Println(distanceFromP(origin))
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())

	scaleP := p.ScaleBy

	scaleP(2)
	scaleP(3)
	scaleP(10)

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	pap := Point{2, 4}
	pptr := &pap
	pptr.ScaleBy(2)
	fmt.Println(pap)

	pop := Point{4, 5}
	(&pop).ScaleBy(2)
	fmt.Println(pop)

}
