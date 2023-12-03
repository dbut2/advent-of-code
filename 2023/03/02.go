package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/dbut2/advent-of-code/pkg/grid"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(2, 467835)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var total int

	g := grid.Grid[*int]{}

	// parse the input and store a pointer to the buffer for each number found
	// all digits in a sequence that make up a number will all point to the same number in grid
	buffer := new(int)
	bufferCells := [][2]int{}
	for i, line := range s {
		for j, char := range line {
			if char >= '0' && char <= '9' {
				*buffer *= 10
				*buffer += int(char - '0')
				bufferCells = append(bufferCells, [2]int{i, j})
			} else {
				for _, cell := range bufferCells {
					g.Set(cell[0], cell[1], buffer)
				}
				buffer = new(int)
				bufferCells = [][2]int{}
			}
		}

		for _, cell := range bufferCells {
			g.Set(cell[0], cell[1], buffer)
		}
		buffer = new(int)
		bufferCells = [][2]int{}
	}

	// find all cogs, find all unique numbers in surround squares, if 2 found then add product
	for i, line := range s {
		for j, char := range line {
			if char == '*' {
				set := sets.Set[*int]{}

				for a := i - 1; a <= i+1; a++ {
					for b := j - 1; b <= j+1; b++ {
						ptr := g.Get(a, b)
						if *ptr != nil {
							set.Add(*ptr)
						}
					}
				}

				if len(set) == 2 {
					sl := set.Slice()
					total += *sl[0] * *sl[1]
				}
			}
		}
	}

	return total
}
