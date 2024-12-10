package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
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
	h.Expect(1, 6)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	a := sets.Set[[2]int]{}
	intersections := [][2]int{}

	for j, line := range s {
		coord := [2]int{0, 0}
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

				switch j {
				case 0:
					a.Add(coord)
				case 1:
					if a.Contains(coord) {
						intersections = append(intersections, coord)
					}
				}
			}
		}
	}

	m := math.MaxInt
	for _, coord := range intersections {
		m = min(m, math.Abs(coord[0])+math.Abs(coord[1]))
	}
	return m
}
