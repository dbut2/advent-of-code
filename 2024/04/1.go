package main

import (
	"slices"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input space.Grid[byte]) int {
	word := "XMAS"

	count := 0
	for cell := range input.Cells() {
		for _, dir := range slices.Concat(space.Directions, space.Diagonals) {
			found := true
			for i := range word {
				nextCell := input.Get(cell.Move(dir.Multiply(i)))
				if nextCell == nil || *nextCell != word[i] {
					found = false
					break
				}
			}
			if found {
				count++
			}
		}
	}

	return count
}

func main() {
	h := harness.New(solve)
	h.Run()
}
