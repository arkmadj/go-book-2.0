package aprint

type format struct {
	name, magic string
	magicOffset int
	reader      NewReader
}
