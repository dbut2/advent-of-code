package space

type Grid[T any] [][]T

func NewGrid[T any](x, y int) Grid[T] {
	g := make([][]T, x)
	for i := 0; i < x; i++ {
		g[i] = make([]T, y)
	}
	return g
}

func NewGridFromInput(s []string) Grid[uint8] {
	g := NewGrid[uint8](len(s[0]), len(s))
	for j := range s {
		for i := range s[j] {
			g[i][j] = s[j][i]
		}
	}
	return g
}

func (g Grid[T]) Inside(x, y int) bool {
	if x < 0 || x >= len(g) {
		return false
	}

	if y < 0 || y >= len(g[0]) {
		return false
	}

	return true
}

func (g Grid[T]) offsets(x, y int, offsets [][2]int) []*T {
	cells := make([]*T, 0, len(offsets))
	for _, coord := range offsets {
		if g.Inside(x+coord[0], y+coord[1]) {
			cells = append(cells, &g[x+coord[0]][y+coord[1]])
		}
	}
	return cells
}

func (g Grid[T]) Adjacent(x, y int) []*T {
	return g.offsets(x, y, [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}})
}

func (g Grid[T]) Diagonal(x, y int) []*T {
	return g.offsets(x, y, [][2]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}})
}

func (g Grid[T]) Surrounding(x, y int) []*T {
	c := make([]*T, 0, 8)
	c = append(c, g.Adjacent(x, y)...)
	c = append(c, g.Diagonal(x, y)...)
	return c
}

func (g Grid[T]) Find(f func(T) bool) *T {
	for i := range g {
		for j, cell := range g[i] {
			if f(cell) {
				return &g[i][j]
			}
		}
	}
	return nil
}
