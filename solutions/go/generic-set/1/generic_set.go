package genericset

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T], 0)
}

func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

func (s Set[T]) Contains(elem T) bool {
	if _, ok := s[elem]; ok {
		return true
	} else {
		return false
	}
}

// Check if two sets have the same elements.
func (s Set[T]) Equals(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for elem := range s {
		if !other.Contains(elem) {
			return false
		}
	}
	for elem := range other {
		if !s.Contains(elem) {
			return false
		}
	}
	return true
}
