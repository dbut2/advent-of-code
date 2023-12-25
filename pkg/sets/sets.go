package sets

import (
	"sync"
)

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
	l := make([]T, 0, len(*s))
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

func (s *Set[T]) Contains(v T) bool {
	b, ok := (*s)[v]
	return ok && b
}

func (s *Set[T]) Copy() Set[T] {
	var c Set[T]
	for i := range *s {
		c.Add(i)
	}
	return c
}

type SyncSet[T comparable] struct {
	set Set[T]
	l   sync.Locker
}

func (s *SyncSet[T]) Slice() []T {
	s.ensureLocker()
	s.l.Lock()
	defer s.l.Unlock()
	return s.set.Slice()
}

func (s *SyncSet[T]) Add(v T) {
	s.ensureLocker()
	s.l.Lock()
	defer s.l.Unlock()
	s.set.Add(v)
}

func (s *SyncSet[T]) Remove(v T) {
	s.ensureLocker()
	s.l.Lock()
	defer s.l.Unlock()
	s.set.Remove(v)
}

func (s *SyncSet[T]) Has(v T) bool {
	s.ensureLocker()
	s.l.Lock()
	defer s.l.Unlock()
	return s.set.Contains(v)
}

func (s *SyncSet[T]) ensureLocker() {
	if s.l == nil {
		s.l = &sync.Mutex{}
	}
}
