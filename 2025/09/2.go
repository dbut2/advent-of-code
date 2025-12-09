package main

import (
	"slices"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
	. "github.com/dbut2/advent-of-code/pkg/std"
)

func solve(input [][]int) int {
	// Find all unique X and Y values, order them and the convert to a smaller
	// index value to create a smaller grid so we're not ranging massive numbers
	xset := sets.Set[int]{}
	yset := sets.Set[int]{}
	for _, line := range input {
		xset.Add(line[0])
		xset.Add(line[0] + 1)
		yset.Add(line[1])
		yset.Add(line[1] + 1)
	}
	xs := xset.Slice()
	ys := yset.Slice()
	slices.Sort(xs)
	slices.Sort(ys)
	shrinkX := make(map[int]int)
	shrinkY := make(map[int]int)
	for i, x := range xs {
		shrinkX[x] = i
	}
	for i, y := range ys {
		shrinkY[y] = i
	}

	// Create our grid, and grid reference of original points
	g := space.NewGrid[bool](len(xs)-1, len(ys)-1)
	gref := space.NewGrid[bool](len(xs)-1, len(ys)-1)
	for _, row := range input {
		sx, sy := shrinkX[row[0]], shrinkY[row[1]]
		g.Set([2]int{sx, sy}, true)
		gref.Set([2]int{sx, sy}, true)
	}

	// Find our start value, we know from where this is positioned the last
	// direction to get to this point was Up
	start := space.Cell{0, 0}
	for y := range ys {
		if v := g.Get(space.Cell{0, y}); v != nil && *v {
			start = space.Cell{0, y}
			break
		}
	}

	// From our start cell and while we are seeing new cells, follow the cells
	// around to create a perimeter around our valid area
	dir := space.Up
	last := start
	seen := sets.Set[space.Cell]{}
	for !seen.Contains(last) {
		g.Set(last, true)
		dir = dir.Rotate().Rotate()
		seen.Add(last)
		for {
			dir = dir.Rotate()
			ray := last
			ray = ray.Move(dir)
			for g.Inside(ray) {
				if v := gref.Get(ray); *v {
					trace := last
					for trace != ray {
						trace = trace.Move(dir)
						g.Set(trace, true)
					}
					last = ray
					goto done
				}
				ray = ray.Move(dir)
			}
		}
	done:
	}

	// Flood fill inside the perimeter to mark all inside calls as valid
	cells := space.FloodCell(g, start.Move(space.Down.Add(space.Right)))
	for _, cell := range cells {
		g.Set(cell, true)
	}

	// Find all pairs of corners from the input and sort by largest area
	var pairs [][2]space.Cell
	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			a, b := input[i], input[j]
			pairs = append(pairs, [2]space.Cell{space.Cell(a), space.Cell(b)})
		}
	}
	slices.SortFunc(pairs, func(a, b [2]space.Cell) int {
		areaA := (Abs(a[0][0]-a[1][0]) + 1) * (Abs(a[0][1]-a[1][1]) + 1)
		areaB := (Abs(b[0][0]-b[1][0]) + 1) * (Abs(b[0][1]-b[1][1]) + 1)
		return areaB - areaA
	})

	// Starting from the largest area, find the 4 edges of each rectangle and
	// return the area of the first rectangle where all edges are entirely
	// inside the valid region
	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		minX := min(shrinkX[a[0]], shrinkX[b[0]])
		maxX := max(shrinkX[a[0]], shrinkX[b[0]])
		minY := min(shrinkY[a[1]], shrinkY[b[1]])
		maxY := max(shrinkY[a[1]], shrinkY[b[1]])

		tl := space.Cell{minX, minY}
		tr := space.Cell{maxX, minY}
		bl := space.Cell{minX, maxY}
		br := space.Cell{maxX, maxY}

		inside := true
		for cell := tl; cell != tr; cell = cell.Move(space.Right) {
			if v := g.Get(cell); v == nil || !*v {
				inside = false
				break
			}
		}
		for cell := tl; cell != bl; cell = cell.Move(space.Down) {
			if v := g.Get(cell); v != nil && !*v {
				inside = false
				break
			}
		}
		for cell := bl; cell != br; cell = cell.Move(space.Right) {
			if v := g.Get(cell); v != nil && !*v {
				inside = false
				break
			}
		}
		for cell := tr; cell != br; cell = cell.Move(space.Down) {
			if v := g.Get(cell); v != nil && !*v {
				inside = false
				break
			}
		}
		if inside {
			return (Abs(a[0]-b[0]) + 1) * (Abs(a[1]-b[1]) + 1)
		}
	}

	return 0
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 24)
	h.Run()
}
