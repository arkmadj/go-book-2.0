package main

import "text/scanner"

type lexer struct {
	scan  scanner.Scanner
	token rune
}
