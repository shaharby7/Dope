package set

type Set[V comparable] struct {
	data map[V]bool
}

type SetOption[V comparable] func(*Set[V]) *Set[V]

func NewSet[V comparable](options ...SetOption[V]) *Set[V] {
	set := &Set[V]{make(map[V]bool)}
	for _, option := range options {
		set = option(set)
	}
	return set
}

func OptionFromSlice[V comparable](initial_slice []V) SetOption[V] {
	return func(s *Set[V]) *Set[V] {
		for _, val := range initial_slice {
			s.Add(val)
		}
		return s
	}
}

func (s *Set[V]) Add(val V) {
	s.data[val] = true
}

func (s *Set[V]) Exists(val V) bool {
	return s.data[val]
}

func (s *Set[V]) Remove(val V) {
	delete(s.data, val)
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
