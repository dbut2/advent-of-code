package utils

type Set[T comparable] map[T]bool

func SetOf[T comparable](s ...T) Set[T] {
	return SetFrom(s)
}

func SetFrom[T comparable](s []T) Set[T] {
	var set Set[T]
	for _, v := range s {
		set[v] = true
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
