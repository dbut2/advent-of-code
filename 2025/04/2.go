package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	. "github.com/dbut2/advent-of-code/pkg/std"
)

func solve(input Grid) int {
	count := 0

	// do while any rolls have been removed
	anyRemoved := true
	for anyRemoved {
		anyRemoved = false
		for cell, v := range input.Cells() {
			if *v != '@' {
				continue
			}
			c := 0
			for _, sur := range input.Surrounding(cell) {
				if sur != nil && *sur == '@' {
					c++
				}
			}
			if c < 4 {
				// remove by replacing value
				*v = '.'
				anyRemoved = true
				count++
			}
		}
	}
	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 43)
	h.Run()
}
