package lists

type DoublyLinked[T any] struct {
	first, last *dllnode[T]
}

func NewDoublyLinked[T any]() DoublyLinked[T] {
	f, l := new(dllnode[T]), new(dllnode[T])
	f.next = l
	l.prev = f
	return DoublyLinked[T]{first: f, last: l}
}

func (d *DoublyLinked[T]) Empty() bool {
	return d.first.next == d.last
}

func (d *DoublyLinked[T]) Append(v T) {
	n := &dllnode[T]{
		value: v,
		prev:  d.last.prev,
		next:  d.last,
	}

	d.last.prev.next = n
	d.last.prev = n
}

func (d *DoublyLinked[T]) TakeFirst() T {
	if d.Empty() {
		panic("cannot take first from empty list")
	}

	n := d.first.next

	d.first.next = n.next
	n.next.prev = d.first

	return n.value
}

func (d *DoublyLinked[T]) TakeLast() T {
	if d.Empty() {
		panic("cannot take last from empty list")
	}

	n := d.last.prev

	d.last.prev = n.prev
	n.prev.next = d.last

	return n.value
}

type dllnode[T any] struct {
	value      T
	prev, next *dllnode[T]
}
