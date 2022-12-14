package grid

type Grid[T any] map[int]map[int]T

func (g *Grid[T]) Set(x, y int, cell T) {
	g.ensureGrid()
	g.ensureRow(x)
	(*g)[x][y] = cell
}

func (g *Grid[T]) Get(x, y int) T {
	g.ensureGrid()
	g.ensureRow(x)
	g.ensureCell(x, y)
	return (*g)[x][y]
}

func (g *Grid[T]) ensureGrid() {
	if g == nil {
		*g = make(Grid[T])
	}
}

func (g *Grid[T]) ensureRow(x int) {
	if _, ok := (*g)[x]; !ok {
		(*g)[x] = make(map[int]T)
	}
}

func (g *Grid[T]) ensureCell(x, y int) {
	if _, ok := (*g)[x][y]; !ok {
		(*g)[x][y] = *new(T)
	}
}
