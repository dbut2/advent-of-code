package grid

type Grid[T any] map[[2]int]*T

func (g Grid[T]) Set(x, y int, cell T) {
	g[[2]int{x, y}] = &cell
}

func (g Grid[T]) Get(x, y int) *T {
	if v, ok := g[[2]int{x, y}]; ok {
		return v
	}
	cell := new(T)
	g[[2]int{x, y}] = cell
	return cell
}
