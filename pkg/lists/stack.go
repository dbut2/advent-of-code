package lists

type Stack[T any] []T

func (s *Stack[T]) Push(items ...T) {
	*s = append(*s, items...)
}

func (s *Stack[T]) Pop() T {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack[T]) Seq(yield func(T) bool) {
	for len(*s) > 0 {
		item := s.Pop()
		if !yield(item) {
			return
		}
	}
}
