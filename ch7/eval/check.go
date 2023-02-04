package eval

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}
