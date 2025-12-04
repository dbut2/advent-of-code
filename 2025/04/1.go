package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	. "github.com/dbut2/advent-of-code/pkg/std"
)

func solve(input Grid) int {
	count := 0
	for cell, v := range input.Cells() {
		if *v != '@' {
			continue
		}
		c := 0
		for _, sur := range input.Surrounding(cell) {
			// outside grid or is a roll
			if sur != nil && *sur == '@' {
				c++
			}
		}
		if c < 4 {
			count++
		}
	}
	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 13)
	h.Run()
}
