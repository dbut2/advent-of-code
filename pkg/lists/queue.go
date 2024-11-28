package lists

type Queue[T any] Linked[T]

func NewQueue[T any]() Queue[T] {
	return Queue[T](NewLinked[T]())
}

func (q *Queue[T]) Push(items ...T) {
	for _, item := range items {
		(*Linked[T])(q).Append(item)
	}
}

func (q *Queue[T]) Pop() T {
	return (*Linked[T])(q).TakeFirst()
}

func (q *Queue[T]) Peek() T {
	return (*Linked[T])(q).first.next.value
}

func (q *Queue[T]) Seq(yield func(T) bool) {
	for !(*Linked[T])(q).Empty() {
		if !yield(q.Pop()) {
			return
		}
	}
}
