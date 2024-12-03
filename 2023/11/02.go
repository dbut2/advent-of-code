package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/grid"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	//h.Expect(1, 0)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	g := grid.Grid[bool]{}

	for i, line := range s {
		for j, char := range line {
			if char == '#' {
				g.Set(i, j, true)
			}
		}
	}

	x1, x2 := g.XRange()
	y1, y2 := g.YRange()

	emptyX := []int{}
	emptyY := []int{}

	for i := x1; i <= x2; i++ {
		containsStar := false
		for j := y1; j <= y2; j++ {
			if g.Get(i, j) != nil {
				containsStar = true
				break
			}
		}

		if !containsStar {
			emptyX = append(emptyX, i)
		}
	}

	for j := y1; j <= y2; j++ {
		containsStar := false
		for i := x1; i <= x2; i++ {
			if g.Get(i, j) != nil {
				containsStar = true
				break
			}
		}

		if !containsStar {
			emptyY = append(emptyY, j)
		}
	}

	total := 0

	for coords1 := range g {
		for coords2 := range g {
			if coords1 == coords2 {
				continue
			}

			x1, x2 := min(coords1[0], coords2[0]), max(coords1[0], coords2[0])
			y1, y2 := min(coords1[1], coords2[1]), max(coords1[1], coords2[1])

			xDistance := x2 - x1 + len(lists.Filter(emptyX, func(i int) bool {
				return x1 < i && i < x2
			}))*999999
			yDistance := y2 - y1 + len(lists.Filter(emptyY, func(i int) bool {
				return y1 < i && i < y2
			}))*999999

			total += xDistance + yDistance
		}
	}

	total /= 2

	return total
}
