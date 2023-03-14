package main

import "testing"

func newIntSets() []IntSet {
	return []IntSet{&BitIntSet{}, NewMapIntSet()}
}

func TestLenZeroInitially(t *testing.T) {
	for _, s := range newIntSets() {
		if s.Len() != 0 {
			t.Errorf("%T.len(): git %d, want 0", s, s.Len())
		}
	}
}

func TestLenAfterAddingElements(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Add(2000)
		if s.Len() != 2 {
			t.Errorf("%T.Len(): got %d, want 2", s, s.Len())
		}
	}
}
