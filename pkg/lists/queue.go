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

func (q *Queue[T]) Peek() T {
	return (*q)[0]
}

func (q *Queue[T]) Seq(yield func(T) bool) {
	for len(*q) > 0 {
		item := q.Pop()
		if !yield(item) {
			return
		}
	}
}
