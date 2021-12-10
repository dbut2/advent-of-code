package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	h := heightMap{}

	for x, line := range s {
		row := []*cell{}
		for y, v := range strings.Split(line, "") {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err.Error())
			}
			c := &cell{
				height:  n,
				x:       x,
				y:       y,
				visited: false,
			}
			row = append(row, c)
		}
		h = append(h, row)
	}

	var basins [][]*cell

	for x, row := range h {
		for y, c := range row {
			u := h.get(x, y-1)
			d := h.get(x, y+1)
			l := h.get(x-1, y)
			r := h.get(x+1, y)

			if c.height < u && c.height < d && c.height < l && c.height < r && !c.visited {
				basin := c.makeBasin(h)
				basins = append(basins, basin)
			}
		}
	}

	var counts []int

	for _, b := range basins {
		counts = append(counts, len(b))
	}

	sort.Ints(counts)

	return counts[len(counts)-1] * counts[len(counts)-2] * counts[len(counts)-3]
}

type heightMap [][]*cell

type cell struct {
	height  int
	x, y    int
	visited bool
}

func (h heightMap) get(x, y int) int {
	if x < 0 || x >= len(h) {
		return 10
	}
	row := h[x]
	if y < 0 || y >= len(row) {
		return 10
	}
	return row[y].height
}

func (h heightMap) getUnvisitedNeighbours(c *cell) []*cell {
	var ns []*cell

	if h.get(c.x, c.y-1) != 10 {
		n := h[c.x][c.y-1]
		if !n.visited && n.height < 9 && n.height > c.height {
			n.visited = true
			ns = append(ns, n)
		}
	}
	if h.get(c.x, c.y+1) != 10 {
		n := h[c.x][c.y+1]
		if !n.visited && n.height < 9 && n.height > c.height {
			n.visited = true
			ns = append(ns, n)
		}
	}
	if h.get(c.x-1, c.y) != 10 {
		n := h[c.x-1][c.y]
		if !n.visited && n.height < 9 && n.height > c.height {
			n.visited = true
			ns = append(ns, n)
		}
	}
	if h.get(c.x+1, c.y) != 10 {
		n := h[c.x+1][c.y]
		if !n.visited && n.height < 9 && n.height > c.height {
			n.visited = true
			ns = append(ns, n)
		}
	}

	return ns
}

func (c *cell) makeBasin(h heightMap) []*cell {
	basin := []*cell{c}
	c.visited = true
	for _, neighbour := range h.getUnvisitedNeighbours(c) {
		newBasin := neighbour.makeBasin(h)
		basin = append(basin, newBasin...)
	}
	return basin
}
