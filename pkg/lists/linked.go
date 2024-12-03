package lists

type Linked[T any] struct {
	first, last *llnode[T]
}

func NewLinked[T any]() Linked[T] {
	first := &llnode[T]{}
	last := &llnode[T]{}
	first.next = last
	return Linked[T]{first: first, last: last}
}

func (l *Linked[T]) Empty() bool {
	return l.first.next == l.last
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
