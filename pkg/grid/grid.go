package grid

import (
	"slices"
)

type Grid[T any] map[[2]int]*T

func (g Grid[T]) Set(x, y int, cell T) {
	g[[2]int{x, y}] = &cell
}

func (g Grid[T]) Get(x, y int) *T {
	if v, ok := g[[2]int{x, y}]; ok {
		return v
	}
	return nil
}

func (g Grid[T]) Adjacent(x, y int) []*T {
	c := make([]*T, 0, 4)
	for _, coord := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		c = append(c, g.Get(x+coord[0], y+coord[1]))
	}
	return c
}

func (g Grid[T]) Diagonal(x, y int) []*T {
	c := make([]*T, 0, 4)
	for _, coord := range [][2]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}} {
		c = append(c, g.Get(x+coord[0], y+coord[1]))
	}
	return c
}

func (g Grid[T]) Surrounding(x, y int) []*T {
	c := make([]*T, 0, 8)
	c = slices.Concat(c, g.Adjacent(x, y), g.Diagonal(x, y))
	return c
}

func (g Grid[T]) XRange() (x1, x2 int) {
	for coord := range g {
		x1 = coord[0]
		x2 = coord[0]
		break
	}
	for coord := range g {
		x1 = min(x1, coord[0])
		x2 = max(x2, coord[0])
	}
	return
}

func (g Grid[T]) YRange() (y1, y2 int) {
	for coord := range g {
		y1 = coord[1]
		y2 = coord[1]
		break
	}
	for coord := range g {
		y1 = min(y1, coord[1])
		y2 = max(y2, coord[1])
	}
	return
}

func (g Grid[T]) InRange(x1, y1, x2, y2 int) (s []*T) {
	for coord, cell := range g {
		if coord[0] < x1 || coord[1] < y1 || coord[0] > x2 || coord[1] > y2 {
			continue
		}
		s = append(s, cell)
	}
	return
}

func (g Grid[T]) Find(f func(T) bool) *T {
	for _, v := range g {
		if f(*v) {
			return v
		}
	}
	return nil
}
