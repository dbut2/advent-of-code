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
