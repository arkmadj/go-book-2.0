package main

import "fmt"

func main() {
	type Point struct {
		X, Y int
	}

	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}

	w := Wheel{Circle{Point{8, 8}, 5}, 20}

	wa := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w)
	w.X = 42
	fmt.Printf("%#v\n", wa)
}
