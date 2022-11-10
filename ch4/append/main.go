package main

import "fmt"

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z[len(x):], y)
	}
	copy(z[len(x):], y)
	// z[len(x)] = y
	return z
}

func main() {
	x := []int{0, 1, 2, 3, 4}
	y := []int{5, 6, 8, 9}
	// for i := 0; i < 10; i++ {
	// 	y = appendInt(x, i)
	// 	fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	// 	x = y
	// }
	a := appendInt(x, y...)
	fmt.Println(a)

	fmt.Println("-----------")
	fmt.Println("-----------")
	var z []int
	z = append(z, 1)
	z = append(z, 2, 3)
	z = append(z, 4, 5, 6)
	z = append(z, z...)
	fmt.Println(z)
}
