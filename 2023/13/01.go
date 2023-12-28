package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/grid"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 405)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var grids []grid.Grid[bool]

	counter := 0
	offsetI := 0
	for i, line := range s {
		if line == "" {
			offsetI = i + 1
			counter++
			continue
		}

		if len(grids) <= counter {
			grids = append(grids, make(grid.Grid[bool]))
		}

		for j, char := range line {
			grids[counter].Set(j, i-offsetI, char == '#')
		}
	}

	total := 0

	for _, g := range grids {
		x1, x2 := g.XRange()
		y1, y2 := g.YRange()

		for tryReflectX := x1 + 1; tryReflectX <= x2; tryReflectX++ {
			allMatches := true
			for i1 := x1; i1 <= x2; i1++ {
				i2 := tryReflectX + (tryReflectX - i1) - 1

				for j := y1; j <= y2; j++ {
					a := g.Get(i1, j)
					b := g.Get(i2, j)
					if a == nil || b == nil {
						continue
					}
					if *a != *b {
						allMatches = false
					}
				}
			}

			if allMatches {
				total += tryReflectX
				break
			}
		}

		for tryReflectY := y1 + 1; tryReflectY <= y2; tryReflectY++ {
			allMatches := true
			for j1 := y1; j1 <= y2; j1++ {
				j2 := tryReflectY + (tryReflectY - j1) - 1

				for i := x1; i <= x2; i++ {
					a := g.Get(i, j1)
					b := g.Get(i, j2)
					if a == nil || b == nil {
						continue
					}
					if *a != *b {
						allMatches = false
					}
				}
			}

			if allMatches {
				total += tryReflectY * 100
				break
			}
		}
	}

	return total
}
