package main

import (
	"slices"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input [][]int) [2]int {
	g := space.NewGrid[bool](71, 71)
	for b := 0; ; b++ {
		g.Set(space.Cell{input[b][0], input[b][1]}, true)
		for _, flood := range space.FloodAll(g) {
			if slices.Contains(flood, space.Cell{0, 0}) && !slices.Contains(flood, space.Cell{70, 70}) {
				return [2]int(input[b])
			}
		}
	}
}

func main() {
	h := harness.New(solve)
	h.Run()
}
