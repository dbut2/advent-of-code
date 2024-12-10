package main

import (
	"time"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input space.Grid[byte]) int {
	locations := map[byte][]space.Cell{}
	for cell, val := range input.Cells() {
		if *val == '.' {
			continue
		}
		locations[*val] = append(locations[*val], cell)
	}

	antinodes := sets.Set[space.Cell]{}
	for _, list := range locations {
		for i := range list {
			for j := range list {
				if i == j {
					continue
				}

				next := list[i].Move(diff(list[i], list[j]))
				if input.Inside(next) {
					antinodes.Add(next)
				}
			}
		}
	}

	return len(antinodes)
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 14)
	h.Run()
}
