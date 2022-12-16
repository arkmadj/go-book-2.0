package main

const wordSize = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}
