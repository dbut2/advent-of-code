package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 31)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")
	grid := Grid{}

	for i := range s {
		var line []*Cell
		for range s[i] {
			line = append(line, &Cell{})
		}
		grid = append(grid, line)
	}

	var start *Cell

	for i, line := range s {
		for j, cell := range line {
			grid[i][j].MinToEnd = -1

			if string(cell) == "S" {
				start = grid[i][j]
				grid[i][j].val = 97
				continue
			}
			if string(cell) == "E" {
				grid[i][j].MinToEnd = 0
				grid[i][j].val = 122
				continue
			}

			grid[i][j].val = int(cell)
		}
	}

	for i, row := range grid {
		for j, cell := range row {
			var p []*Cell

			if i > 0 {
				p = append(p, grid[i-1][j])
			}
			if i < len(grid)-1 {
				p = append(p, grid[i+1][j])
			}
			if j > 0 {
				p = append(p, grid[i][j-1])
			}
			if j < len(grid[i])-1 {
				p = append(p, grid[i][j+1])
			}

			for _, n := range p {
				if n.val-cell.val <= 1 {
					cell.Neighbours = append(cell.Neighbours, n)
				}
			}
		}
	}

	return FindPath(grid, start)
}

type Grid [][]*Cell

type Cell struct {
	val        int
	Neighbours []*Cell
	MinToEnd   int
}

func FindPath(grid Grid, start *Cell) int {
	for changed := true; changed; {
		changed = false

		for _, row := range grid {
			for _, cell := range row {
				for _, n := range cell.Neighbours {

					if n.MinToEnd == -1 {
						continue
					}

					cm := n.MinToEnd + 1

					if cm < cell.MinToEnd || cell.MinToEnd == -1 {
						cell.MinToEnd = cm
						changed = true
					}
				}

			}
		}

		_ = func() string { return "" }
	}

	return start.MinToEnd
}
