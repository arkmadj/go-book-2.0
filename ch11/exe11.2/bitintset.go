package main

type BitIntSet struct {
	words []uint64
}

func (s *BitIntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *BitIntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *BitIntSet) AddAll(nums ...int) {
	for _, n := range nums {
		s.Add(n)
	}
}

func (s *BitIntSet) UnionWith(t IntSet) {
	if bis, ok := t.(*BitIntSet); ok {
		for i, tword := range bis.words {
			if i < len(s.words) {
				s.words[i] |= tword
			} else {
				s.words = append(s.words, tword)
			}
		}
	} else {
		for _, i := range t.Ints() {
			s.Add(i)
		}
	}
}
