package main

import (
	"embed"
	"slices"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 30)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	a := map[[2]int]int{}
	intersections := []int{}

	for j, line := range s {
		coord := [2]int{0, 0}
		distance := 0
		for _, move := range strings.Split(line, ",") {
			for i := 0; i < sti.Sti(move[1:]); i++ {
				switch move[0] {
				case 'U':
					coord[1]--
				case 'D':
					coord[1]++
				case 'L':
					coord[0]--
				case 'R':
					coord[0]++
				}
				distance++

				switch j {
				case 0:
					if _, ok := a[coord]; !ok {
						a[coord] = distance
					}
				case 1:
					if aDistance, ok := a[coord]; ok {
						intersections = append(intersections, aDistance+distance)
					}
				}
			}
		}
	}

	return slices.Min(intersections)
}
