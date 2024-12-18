package main

import (
	"embed"
	_ "embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 4361)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	grid := space.Grid[*int]{}

	// parse the input and store a pointer to the buffer for each number found
	// all digits in a sequence that make up a number will all point to the same number in grid
	buffer := new(int)
	bufferCells := space.Cells{}
	for j, line := range s {
		for i, char := range line {
			if char >= '0' && char <= '9' {
				*buffer *= 10
				*buffer += int(char - '0')
				bufferCells = append(bufferCells, space.Cell{i, j})
			} else {
				for _, cell := range bufferCells {
					grid.Set(cell, buffer)
				}
				buffer = new(int)
				bufferCells = space.Cells{}
			}
		}

		for _, cell := range bufferCells {
			grid.Set(cell, buffer)
		}
		buffer = new(int)
		bufferCells = space.Cells{}
	}

	numbers := sets.Set[*int]{}
	grid2 := space.NewGridFromInput(s)
	for cell, char := range grid2.Cells() {
		if (*char == '.') || (*char >= '0' && *char <= '9') {
			continue
		}

		for _, neighbourCell := range grid.Surrounding(cell) {
			if *neighbourCell == nil {
				continue
			}
			numbers.Add(*neighbourCell)
		}
	}

	total := 0
	for num := range numbers {
		total += *num
	}
	return total
}
