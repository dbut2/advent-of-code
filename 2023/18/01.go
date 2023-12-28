package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/algorithms"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/space"
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

	coords := space.Cells{}
	coord := space.Cell{0, 0}
	lineLength := 0

	for _, line := range s {
		splits := strings.Split(line, " ")
		direction := splits[0]
		amount := sti.Sti(splits[1])
		switch direction {
		case "U":
			coord[1] += amount
		case "D":
			coord[1] -= amount
		case "L":
			coord[0] -= amount
		case "R":
			coord[0] += amount
		}
		coords = append(coords, coord)
		lineLength += amount
	}

	return algorithms.Shoelace(coords) + lineLength/2 + 1
}
