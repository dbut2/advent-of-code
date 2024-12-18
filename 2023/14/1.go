package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 136)
	h.Run()
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
