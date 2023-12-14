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
	h.Expect(1, 136)
	h.Solve()
}

type rock int

const (
	None rock = iota
	Round
	Cube
)

func solve(input string) int {
	s := utils.ParseInput(input)

	g := grid.Grid[rock]{}

	for j, line := range s {
		for i, char := range line {
			switch char {
			case '.':
				g.Set(i, j, None)
			case 'O':
				g.Set(i, j, Round)
			case '#':
				g.Set(i, j, Cube)
			}
		}
	}

	x1, x2 := g.XRange()
	y1, y2 := g.YRange()

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if *g.Get(i, j) != Round {
				continue
			}

			j2 := j
			for j2 > 0 && *g.Get(i, j2-1) == None {
				j2 -= 1
			}

			g.Set(i, j, None)
			g.Set(i, j2, Round)
		}
	}

	total := 0
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if *g.Get(i, j) == Round {
				total += y2 - j + 1
			}
		}
	}

	return total
}
