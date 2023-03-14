package main

func newIntSets() []IntSet {
	return []IntSet{&BitIntSet{}, NewMapIntSet()}
}
