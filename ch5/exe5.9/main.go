package main

import (
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`\$\w+|\${\w+}`)

func expand(s string, f func(string) string) string {
	wrapper := func(s string) string {
		if strings.HasPrefix(s, "${") {
			s = s[2 : len(s)-1]
		} else {
			s = s[1:]
		}
		return f(s)
	}
	return pattern.ReplaceAllStringFunc(s, wrapper)
}
