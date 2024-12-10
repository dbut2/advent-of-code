package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/graphs"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 102)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	grid := space.NewGridFromInput(s)

	type node struct {
		cell  space.Cell
		layer int
	}
	graph := graphs.New[node]()

	start := node{cell: space.Cell{0, 0}}
	end := node{cell: space.Cell{len(grid) - 1, len(grid[0]) - 1}}

	for cell := range grid.Cells() {
		for _, direction := range []space.Direction{space.Up, space.Down, space.Left, space.Right} {
			nextCell := cell
			distance := 0
			for i := 1; i <= 3; i++ {
				nextCell = nextCell.Move(direction)
				if !grid.Inside(nextCell) {
					break
				}
				nextC := grid.Get(nextCell)
				distance += int(*nextC - '0')

				a := node{cell, math.Abs(direction[0])}
				b := node{nextCell, math.Abs(direction[1])}
				if a.cell == start.cell || a.cell == end.cell {
					a.layer = 0
				}
				if b.cell == start.cell || b.cell == end.cell {
					b.layer = 0
				}
				graph.Connect(a, b, distance)
			}
		}
	}

	return graph.Minimise(start, end)
}
