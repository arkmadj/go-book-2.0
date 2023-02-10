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

func (lex *lexer) next()
