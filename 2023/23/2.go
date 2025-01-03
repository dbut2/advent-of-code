package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/graphs"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 154)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	grid := space.NewGridFromInput(s)

	start, _ := grid.Find(func(cell space.Cell, char uint8) bool {
		return cell[1] == 0 && char == '.'
	})
	end, _ := grid.Find(func(cell space.Cell, char uint8) bool {
		return cell[1] == len(grid[0])-1 && char == '.'
	})

	graph := graphs.Graph[space.Cell]{}

	traversed := sets.Set[[2]space.Cell]{}

	type job struct {
		lastNode space.Cell
		lastCell space.Cell
		thisCell space.Cell
		length   int
	}
	current := lists.NewQueue[job]()
	current.Push(job{
		lastNode: start,
		lastCell: start,
		thisCell: start.Move(space.Down),
		length:   1,
	})
	for j := range current.Seq {
		// dont re-traverse a path
		if traversed.Contains([2]space.Cell{j.lastCell, j.thisCell}) {
			continue
		}
		traversed.Add([2]space.Cell{j.lastCell, j.thisCell})

		// add end cell as node
		if j.thisCell == end {
			graph.Connect(j.lastNode, j.thisCell, j.length)
			continue
		}

		// pull adjacent paths
		adjacent := grid.Adjacent(j.thisCell)
		for k, v := range adjacent {
			if *v == '#' {
				delete(adjacent, k)
			}
		}

		// determine if node
		if len(adjacent) > 2 {
			graph.Connect(j.lastNode, j.thisCell, j.length)
			j.lastNode = j.thisCell
			j.length = 0
		}

		// don't double back
		for k := range adjacent {
			if k == j.lastCell {
				delete(adjacent, k)
				continue
			}
		}

		// traverse next paths
		for k := range adjacent {
			current.Push(job{
				lastNode: j.lastNode,
				lastCell: j.thisCell,
				thisCell: k,
				length:   j.length + 1,
			})
		}
	}

	return graph.Maximise(start, end)
}
