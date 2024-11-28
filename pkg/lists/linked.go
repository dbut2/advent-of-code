package lists

type Linked[T any] struct {
	first, last *llnode[T]
}

func NewLinked[T any]() Linked[T] {
	n := &llnode[T]{}
	return Linked[T]{first: n, last: n}
}

func (l *Linked[T]) Empty() bool {
	return l.first.next == nil
}

func (l *Linked[T]) Append(v T) {
	l.last.next = &llnode[T]{value: v}
	l.last = l.last.next
}

func (l *Linked[T]) TakeFirst() T {
	n := l.first.next
	l.first.next = l.first.next.next
	return n.value
}

type llnode[T any] struct {
	value T
	next  *llnode[T]
}
