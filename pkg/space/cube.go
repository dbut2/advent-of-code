package space

import (
	"slices"
)

type Cube[T any] [][][]T

func NewCube[T any](x, y, z int) Cube[T] {
	c := make([][][]T, x)
	for i := range x {
		c[i] = make([][]T, y)
		for j := 0; j < y; j++ {
			c[i][j] = make([]T, z)
		}
	}
	return c
}

func (c Cube[T]) Inside(x, y, z int) bool {
	if x < 0 || x >= len(c) {
		return false
	}

	if y < 0 || y >= len(c[0]) {
		return false
	}

	if z < 0 || z >= len(c[0][0]) {
		return false
	}

	return true
}

func (c Cube[T]) offsets(x, y, z int, offsets [][3]int) []*T {
	cells := make([]*T, 0, len(offsets))
	for _, coord := range offsets {
		if c.Inside(x+coord[0], y+coord[1], z+coord[2]) {
			cells = append(cells, &c[x+coord[0]][y+coord[1]][z+coord[2]])
		}
	}
	return cells
}

func (c Cube[T]) Adjacent(x, y, z int) []*T {
	return c.offsets(x, y, z, [][3]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}})
}

func (c Cube[T]) Edges(x, y, z int) []*T {
	return c.offsets(x, y, z, [][3]int{{1, 1, 0}, {-1, 1, 0}, {1, -1, 0}, {-1, -1, 0}, {1, 0, 1}, {-1, 0, 1}, {1, 0, -1}, {-1, 0, -1}, {0, 1, 1}, {0, -1, 1}, {0, 1, -1}, {0, -1, -1}})
}

func (c Cube[T]) Corners(x, y, z int) []*T {
	return c.offsets(x, y, z, [][3]int{{1, 1, 1}, {-1, 1, 1}, {1, -1, 1}, {-1, -1, 1}, {1, 1, -1}, {-1, 1, -1}, {1, -1, -1}, {-1, -1, -1}})
}

func (c Cube[T]) Surrounding(x, y, z int) []*T {
	cells := make([]*T, 0, 26)
	cells = slices.Concat(cells, c.Adjacent(x, y, z), c.Edges(x, y, z), c.Corners(x, y, z))
	return cells
}

func (c Cube[T]) Find(f func(T) bool) *T {
	for i := range c {
		for j := range c[i] {
			for k, cell := range c[i][j] {
				if f(cell) {
					return &c[i][j][k]
				}
			}
		}
	}
	return nil
}
