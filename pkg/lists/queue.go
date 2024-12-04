package lists

type Queue[T any] []T

func NewQueue[T any]() Queue[T] {
	return make(Queue[T], 0)
}

func (q *Queue[T]) Push(items ...T) {
	*q = append(*q, items...)
}

func (q *Queue[T]) Pop() T {
	item := q.Peek()
	*q = (*q)[:len(*q)-1]
	return item
}

func (q *Queue[T]) Peek() T {
	return (*q)[len(*q)-1]
}

func (q *Queue[T]) Seq(yield func(T) bool) {
	for len(*q) > 0 {
		if !yield(q.Pop()) {
			return
		}
	}
}

// TODO: fix the failing cases when using ll
//type Queue[T any] struct {
//	linked Linked[T]
//}

//func NewQueue[T any]() Queue[T] {
//	return Queue[T]{linked: NewLinked[T]()}
//}
//
//func (q *Queue[T]) Push(items ...T) {
//	for _, item := range items {
//		q.linked.Append(item)
//	}
//}
//
//func (q *Queue[T]) Pop() T {
//	return q.linked.TakeFirst()
//}
//
//func (q *Queue[T]) Peek() T {
//	return q.linked.first.next.value
//}
//
//func (q *Queue[T]) Seq(yield func(T) bool) {
//	for !q.linked.Empty() {
//		if !yield(q.Pop()) {
//			return
//		}
//	}
//}
