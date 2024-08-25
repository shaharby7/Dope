package utils

type Set[V comparable] struct {
	data map[V]bool
}

func NewSet[V comparable]() *Set[V] {
	return &Set[V]{make(map[V]bool)}
}

func (s *Set[V]) Add(val V) {
	s.data[val] = true
}

func (s *Set[V]) Exists(val V) bool {
	return s.data[val]
}

func (s *Set[V]) Size() int {
	return len(s.data)
}

func (s *Set[V]) ToSlice() []V {
	i := 0
	keys := make([]V, len(s.data))
	for k := range s.data {
		keys[i] = k
		i++
	}
	return keys
}
