package sets

type Set[T comparable] map[T]bool

func SetOf[T comparable](s ...T) Set[T] {
	return SetFrom(s)
}

func SetFrom[T comparable](s []T) Set[T] {
	set := make(Set[T])
	for _, v := range s {
		set.Add(v)
	}
	return set
}

func (s *Set[T]) Slice() []T {
	var l []T
	for i := range *s {
		l = append(l, i)
	}
	return l
}

func (s *Set[T]) Add(v T) {
	if *s == nil {
		*s = make(map[T]bool)
	}
	(*s)[v] = true
}

func (s *Set[T]) Remove(v T) {
	delete(*s, v)
}

func (s *Set[T]) Has(v T) bool {
	_, ok := (*s)[v]
	return ok
}