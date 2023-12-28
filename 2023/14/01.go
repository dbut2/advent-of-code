package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/space"
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

func solve(input string) int {
	s := utils.ParseInput(input)

	grid := space.NewGridFromInput(s)

	total := 0
	for _, column := range grid {
		topFree := 0
		for j, cell := range column {
			switch cell {
			case '#':
				topFree = j + 1
			case 'O':
				total += len(s) - topFree
				topFree++
			}
		}
	}
	return total
}
