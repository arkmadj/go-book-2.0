package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
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

type lexPanic string

type selector struct {
	tag   string
	attrs []attribute
}

func (s selector) String() string {
	b := &bytes.Buffer{}
	b.WriteString(s.tag)
	for _, attr := range s.attrs {
		switch attr.Value {
		case "":
			fmt.Fprintf(b, "[%s]", attr.Name)
		default:
			fmt.Fprintf(b, `[%s="%s"]`, attr.Name, attr.Value)
		}
	}
	return b.String()
}

type attribute struct {
	Name, Value string
}

func attrMatch(selAttrs []attribute, xmlAttrs []xml.Attr) bool {
SelectorAttribute:
	for _, sa := range selAttrs {
		for _, xa := range xmlAttrs {
			if sa.Name == xa.Name.Local && sa.Value == xa.Value || sa.Value == "" {
				continue SelectorAttribute
			}
		}
		return false
	}
	return true
}

func parseSelectors(input string) (_ []selector, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			panic(x)
		}
	}()

	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanStrings
	lex.scan.Whitespace = 0
	lex.next()

	selectors := make([]selector, 0)
	for lex.token != scanner.EOF {
		selectors = append(selectors, parseSelector(lex))
	}
	return selectors, nil
}

func parseSelector(lex *lexer) selector {
	var sel selector
	lex.eatWhitespace()
	if lex.token != '[' {
		if lex.token != scanner.Ident {
			panic(lexPanic(fmt.Sprintf("got %s, want ident", lex.describe())))
		}
		sel.tag = lex.text()
		lex.next()
	}
	for lex.token == '[' {
		sel.attrs = append(sel.attrs, parseAttr(lex))
	}
	return sel
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	sels, err := parseSelectors(strings.Join(os.Args[2:], " "))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	xmlselect(os.Stdout, os.Stdin, sels)
}
