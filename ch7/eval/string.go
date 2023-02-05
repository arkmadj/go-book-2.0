package eval

import "fmt"

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.x.String(), string(b.op), b.y.String())
}
