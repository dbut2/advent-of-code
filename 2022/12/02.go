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
	t.Expect(1, 29)
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

	var starts []*Cell

	for i, line := range s {
		for j, cell := range line {
			grid[i][j].MinToEnd = -1

			if string(cell) == "S" {
				starts = append(starts, grid[i][j])
				grid[i][j].val = 97
				continue
			}
			if string(cell) == "E" {
				grid[i][j].MinToEnd = 0
				grid[i][j].val = 122
				continue
			}

			grid[i][j].val = int(cell)

			if grid[i][j].val == 97 {
				starts = append(starts, grid[i][j])
			}
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

	return FindPath(grid, starts)
}

type Grid [][]*Cell

type Cell struct {
	val        int
	Neighbours []*Cell
	MinToEnd   int
}

func FindPath(grid Grid, starts []*Cell) int {
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
	}

	min := -1
	for _, start := range starts {
		if min == -1 || start.MinToEnd < min && start.MinToEnd != -1 {
			min = start.MinToEnd
		}

	}
	return min
}
