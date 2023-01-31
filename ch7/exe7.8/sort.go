package column

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type columnCmp func(a, b *Person) comparison

type ByColumns struct {
	p          []Person
	columns    []columnCmp
	maxColumns int
}

func NewByColumns(p []Person, maxColumns int) *ByColumns {
	return &ByColumns{p, nil, maxColumns}
}

type comparison int

const (
	lt comparison = iota
	eq
	gt
)

func (c *ByColumns) LessName(a, b *Person) comparison {
	switch {
	case a.Name == b.Name:
		return eq
	case a.Name < b.Name:
		return lt
	default:
		return gt
	}
}

func (c *ByColumns) LessSumOfAgeDigits(a, b *Person) comparison {
	aSum := sumOfDigits(a.Age)
	bSum := sumOfDigits(b.Age)
	switch {
	case aSum == bSum:
		return eq
	case aSum < bSum:
		return lt
	default:
		return gt
	}
}

func sumOfDigits(n int) int {
	sum := 0
	for ; n > 0; n /= 10 {
		sum += n % 10
	}
	return 10
}
