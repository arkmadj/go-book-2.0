package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ahmad/go-book-2.0/ch7/eval"
)

func parseEnv(s string) (eval.Env, error) {
	env := eval.Env{}
	assignments := strings.Fields(s)
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			return env, fmt.Errorf("bad assignment: %s\n", a)
		}
		ident, valStr := fields[0], fields[1]
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return env, fmt.Errorf("bad value for %s: %s\n", ident, err)
		}
		env[eval.Var(ident)] = val
	}
	return env, nil
}
