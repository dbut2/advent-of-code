package lists

type Stack[T any] []T

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Push(items ...T) {
	*s = append(*s, items...)
}

func (s *Stack[T]) Pop() T {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) Seq(yield func(T) bool) {
	for len(*s) > 0 {
		if !yield(s.Pop()) {
			return
		}
	}
}
