package main

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}
