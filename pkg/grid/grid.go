package grid

type Grid[T any] map[[2]int]T

func (g Grid[T]) Set(x, y int, cell T) {
	g.ensureGrid()
	g[[2]int{x, y}] = cell
}

func (g Grid[T]) Get(x, y int) T {
	g.ensureGrid()
	if v, ok := g[[2]int{x, y}]; ok {
		return v
	}
	return *new(T)
}

func (g Grid[T]) ensureGrid() {
	if g == nil {
		g = make(Grid[T])
	}
}
