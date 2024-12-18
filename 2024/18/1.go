package main

import (
	"github.com/dbut2/advent-of-code/pkg/graphs"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input [][]int) int {
	g := space.NewGrid[bool](71, 71)
	graph := graphs.New[space.Cell]()
	for _, line := range input[:1024] {
		g.Set(space.Cell{line[0], line[1]}, true)
	}
	for cell, v := range g.Cells() {
		for next, w := range g.Adjacent(cell) {
			if !*v && !*w {
				graph.Connect(cell, next, 1)
			}
		}
	}
	return graph.Minimise(space.Cell{0, 0}, space.Cell{70, 70})
}

func main() {
	h := harness.New(solve)
	h.Run()
}
