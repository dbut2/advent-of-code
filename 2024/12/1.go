package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input space.Grid[byte]) int {
	floods := space.FloodAll(input)

	total := 0
	for _, flood := range floods {
		perimeter := 0
		for _, cell := range flood {
			for _, dir := range space.Directions {
				next := input.Get(cell.Move(dir))
				if next == nil || *next != *input.Get(cell) {
					perimeter++
				}
			}
		}
		total += perimeter * len(flood)
	}

	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 772)
	h.Expect(2, 1930)
	h.Run()
}
