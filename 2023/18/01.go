package main

import (
	"embed"
	"math"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 62)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	coords := [2]int{0, 0}
	path := sets.Set[[2]int]{}

	for j, line := range s {
		_, _ = j, line

		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")

		parts := strings.Split(line, " ")

		amount := sti.Sti(parts[1])

		switch parts[0] {
		case "U":
			for i := 0; i < amount; i++ {
				coords[0]--
				path.Add(coords)
			}
		case "D":
			for i := 0; i < amount; i++ {
				coords[0]++
				path.Add(coords)
			}
		case "L":
			for i := 0; i < amount; i++ {
				coords[1]--
				path.Add(coords)
			}
		case "R":
			for i := 0; i < amount; i++ {
				coords[1]++
				path.Add(coords)
			}
		}
	}

	x1, x2 := math.MaxInt, math.MinInt
	y1, y2 := math.MaxInt, math.MinInt

	for coord := range path {
		x1, x2 = min(x1, coord[0]), max(x2, coord[0])
		y1, y2 = min(y1, coord[1]), max(y2, coord[1])
	}

	queue := lists.Queue[[2]int]{}
	for x := x1; x <= x2; x++ {
		queue.Push([2]int{x, y1})
		queue.Push([2]int{x, y2})
	}
	for y := y1; y <= y2; y++ {
		queue.Push([2]int{x1, y})
		queue.Push([2]int{x2, y})
	}

	outside := sets.Set[[2]int]{}

	for len(queue) > 0 {
		item := queue.Pop()

		if item[0] < x1 || item[0] > x2 {
			continue
		}
		if item[1] < y1 || item[1] > y2 {
			continue
		}

		if path.Has(item) {
			continue
		}

		if outside.Has(item) {
			continue
		}

		outside.Add(item)

		for _, delta := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			queue.Push([2]int{item[0] + delta[0], item[1] + delta[1]})
		}
	}

	return (x2-x1+1)*(y2-y1+1) - len(outside)
}
