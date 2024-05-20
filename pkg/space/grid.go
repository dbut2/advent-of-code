package space

import (
	"slices"
)

type Grid[T any] [][]T

func NewGrid[T any](x, y int) Grid[T] {
	g := make([][]T, x)
	for i := range x {
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

func (g *Grid[T]) Get(c Cell) *T {
	if !g.Inside(c) {
		return nil
	}
	return &(*g)[c[0]][c[1]]
}

func (g *Grid[T]) Set(c Cell, v T) {
	if !g.Inside(c) {
		g.growTo(c)
	}
	(*g)[c[0]][c[1]] = v
}

func (g *Grid[T]) growTo(c Cell) {
	if g.Inside(c) {
		return
	}

	if c[0] >= len(*g) {
		l := c[1]
		for i := range *g {
			l = max(l, len((*g)[i]))
		}

		for i := len(*g); i <= c[0]; i++ {
			*g = append(*g, make([]T, l))
		}
	}

	if c[1] >= len((*g)[0]) {
		for i := range *g {
			(*g)[i] = slices.Concat((*g)[i], make([]T, c[1]-len((*g)[i])+1))
		}
	}
}

func (g *Grid[T]) Inside(c Cell) bool {
	if c[0] < 0 || c[0] >= len(*g) {
		return false
	}

	if c[1] < 0 || c[1] >= len((*g)[0]) {
		return false
	}

	return true
}

func (g *Grid[T]) offsets(c Cell, offsets []Direction) map[Cell]*T {
	cells := make(map[Cell]*T, len(offsets))
	for _, direction := range offsets {
		next := c.Move(direction)
		if g.Inside(next) {
			cells[next] = &(*g)[next[0]][next[1]]
		}
	}
	return cells
}

func (g *Grid[T]) Adjacent(c Cell) map[Cell]*T {
	return g.offsets(c, []Direction{North, South, East, West})
}

func (g *Grid[T]) Diagonal(c Cell) map[Cell]*T {
	return g.offsets(c, []Direction{North.Add(East), North.Add(West), South.Add(East), South.Add(West)})
}

func (g *Grid[T]) Surrounding(c Cell) map[Cell]*T {
	cells := make(map[Cell]*T, 8)
	for k, v := range g.Adjacent(c) {
		cells[k] = v
	}
	for k, v := range g.Diagonal(c) {
		cells[k] = v
	}
	return cells
}

func (g *Grid[T]) Find(f func(Cell, T) bool) (Cell, *T) {
	for i := range *g {
		for j, cell := range (*g)[i] {
			c := Cell{i, j}
			if f(c, cell) {
				return c, &(*g)[i][j]
			}
		}
	}
	return Cell{}, nil
}

func (g *Grid[T]) Cells() map[Cell]*T {
	cells := make(map[Cell]*T, len(*g)*len((*g)[0]))
	for i := range *g {
		for j := range (*g)[i] {
			cells[Cell{i, j}] = &(*g)[i][j]
		}
	}
	return cells
}
