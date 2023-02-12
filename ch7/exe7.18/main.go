package main

import "encoding/xml"

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

// func (n *Element) String() string {
// 	b := &bytes.Buffer{}
// 	v
// }
