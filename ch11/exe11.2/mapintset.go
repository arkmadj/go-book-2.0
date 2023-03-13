package main

type MapIntSet struct {
	m map[int]bool
}

func NewMapIntSet() *MapIntSet {
	return &MapIntSet{map[int]bool{}}
}

func (s *MapIntSet) Has(x int) bool {
	return s.m[x]
}

func (s *MapIntSet) Add(x int) {
	s.m[x] = true
}

func (s *MapIntSet) AddAll(nums ...int) {
	for _, x := range nums {
		s.m[x] = true
	}
}

func (s *MapIntSet) UnionWith(t IntSet) {
	for _, x := range t.Ints() {
		s.m[x] = true
	}
}

func (s *MapIntSet) Len() int {
	return len(s.m)
}

func (s *MapIntSet) Remove(x int) {
	delete(s.m, x)
}

func (s *MapIntSet) Clear() {
	s.m = make(map[int]bool)
}
