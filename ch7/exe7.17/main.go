package main

import (
	"fmt"
	"text/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token))
}

func (lex *lexer) eatWhitespace() int {
	i := 0
	for lex.token == ' ' || lex.token == '\t' {
		lex.next()
		i++
	}
	return i
}

func (lex *lexer) next() {
	lex.token = lex.scan.Scan()
}
func (lex *lexer) text() string {
	return lex.scan.TokenText()
}
