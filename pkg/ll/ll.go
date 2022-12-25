package ll

type Double[T any] struct {
	Val  T
	Prev *Double[T]
	Next *Double[T]
}

func Find[T comparable](some *Double[T], val T) *Double[T] {
	c := some
	for c.Val != val {
		c = c.Next
	}
	return c
}

func (d *Double[T]) MoveLeft() {
	p0 := d.Prev
	p1 := p0.Prev
	n := d.Next

	d.Prev = p1
	d.Next = p0

	p0.Prev = d
	p0.Next = n

	n.Prev = p0
	p1.Next = d
}

func (d *Double[T]) MoveRight() {
	n0 := d.Next
	n1 := n0.Next
	p := d.Prev

	d.Prev = n0
	d.Next = n1

	n0.Prev = p
	n0.Next = d

	p.Next = n0
	n1.Prev = d
}

func Link[T any](first, second *Double[T]) {
	if first == nil || second == nil {
		return
	}
	first.Next = second
	second.Prev = first
}
