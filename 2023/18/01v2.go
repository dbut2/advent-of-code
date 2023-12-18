package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
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
	var coords [][2]int
	coord := [2]int{0, 0}
	length := 0
	for _, line := range s {
		parts := strings.Split(line, " ")
		amount := sti.Sti(parts[1])
		switch parts[0] {
		case "U":
			coord[1] -= amount
		case "D":
			coord[1] += amount
		case "L":
			coord[0] -= amount
		case "R":
			coord[0] += amount
		}
		coords = append(coords, coord)
		length += amount
	}

	return math.Shoelace(coords) + length/2 + 1
}
