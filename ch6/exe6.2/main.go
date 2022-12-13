package main

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(nums ...int) {
	for _, n := range nums {
		s.Add(n)
	}
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func popcount(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popcount(word)
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &^= 1 << uint64(bit)
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}
