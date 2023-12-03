package buffers

type Number int

func (b *Number) Add(n int) {
	*b *= 10
	*b += Number(n)
}

func (b *Number) Clear() int {
	val := int(*b)
	*b = 0
	return val
}
