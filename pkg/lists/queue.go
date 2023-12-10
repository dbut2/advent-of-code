package lists

type Queue[T any] []T

func (q *Queue[T]) Push(items ...T) {
	*q = append(*q, items...)
}

func (q *Queue[T]) Pop() T {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}
