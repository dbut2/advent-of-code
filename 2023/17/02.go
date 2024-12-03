package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/graphs"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Tester.Expect(1, 94)
	h.Tester.Expect(2, 71)
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
			for i := 1; i <= 10; i++ {
				nextCell = nextCell.Move(direction)
				if !grid.Inside(nextCell) {
					break
				}
				nextC := grid.Get(nextCell)
				distance += int(*nextC - '0')

				if i >= 4 {
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
	}

	return graph.Minimise(start, end)
}
