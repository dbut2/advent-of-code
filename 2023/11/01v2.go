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
	h.Expect(1, 374)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	g := grid.Grid[bool]{}

	for i, line := range s {
		for j, char := range line {
			g.Set(i, j, char == '#')
		}
	}

	x1, x2 := g.XRange()
	y1, y2 := g.YRange()

	emptyX := lists.Filter(lists.Range(x1, x2), func(i int) bool {
		return len(lists.Filter(lists.Range(y1, y2), func(j int) bool {
			return *g.Get(i, j)
		})) == 0
	})
	emptyY := lists.Filter(lists.Range(x1, x2), func(j int) bool {
		return len(lists.Filter(lists.Range(y1, y2), func(i int) bool {
			return *g.Get(i, j)
		})) == 0
	})

	stars := lists.Filter(lists.MapToSlice(g), func(pair lists.Pair[[2]int, *bool]) bool {
		return *pair.Val
	}).Keys()

	total := lists.Reduce(stars, func(total int, star1 [2]int) int {
		return total + lists.Reduce(stars, func(total int, star2 [2]int) int {
			x1, x2 := min(star1[0], star2[0]), max(star1[0], star2[0])
			y1, y2 := min(star1[1], star2[1]), max(star1[1], star2[1])

			xDistance := x2 - x1
			yDistance := y2 - y1

			xDistance += len(lists.Filter(emptyX, func(i int) bool {
				return x1 < i && i < x2
			}))
			yDistance += len(lists.Filter(emptyY, func(i int) bool {
				return y1 < i && i < y2
			}))

			return total + xDistance + yDistance
		})
	})

	return total / 2
}
