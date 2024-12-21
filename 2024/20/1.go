package main

import (
	"github.com/dbut2/advent-of-code/pkg/graphs"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/std"
)

func solve(input space.Grid[byte]) int {
	// Create a graph of all connected cells in the path
	graph := graphs.New[space.Cell]()
	for cell, val := range input.Cells() {
		if *val == '#' {
			continue
		}
		for _, dir := range space.Directions {
			next := cell.Move(dir)
			v := input.Get(next)
			if v == nil || *v == '#' {
				continue
			}
			graph.Connect(cell, next, 1)
		}
	}

	start := input.Find(std.Equal[byte]('S'))
	end := input.Find(std.Equal[byte]('E'))

	base := graph.Minimise(start, end)

	count := 0
	for cell, val := range input.Cells() {
		const k = 2

		if *val == '#' {
			continue
		}

		// Move out up to k distance from cell. Add connection on graph and find
		// new race score using Dijkstra.

		options := sets.SetOf(cell)
		nextOptions := sets.Set[space.Cell]{}
		for range k {
			for c := range options.Seq {
				for _, dir := range space.Directions {
					nextOptions.Add(c.Move(dir))
				}
			}
			options, nextOptions = nextOptions, options
		}

		for c := range options.Seq {
			next := input.Get(c)
			if next == nil || *next == '#' {
				continue
			}
			graph.Connect(cell, c, k)
			l := graph.Minimise(start, end)
			delete(graph.Edges[cell], c)

			if l == -1 {
				continue
			}
			if base-l >= 100 {
				count++
			}
		}
	}
	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 0)
	h.Run()
}
