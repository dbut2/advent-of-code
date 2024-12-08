package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input space.Grid[byte]) int {
	antinodes := sets.Set[space.Cell]{}
	locations := map[byte][]space.Cell{}
	for cell, val := range input.Cells() {
		if *val == '.' {
			continue
		}
		antinodes.Add(cell)
		locations[*val] = append(locations[*val], cell)
	}

	for _, list := range locations {
		for i := range list {
			for j := range list {
				if i == j {
					continue
				}

				next := list[i].Move(diff(list[i], list[j]))
				for input.Inside(next) {
					antinodes.Add(next)
					next = next.Move(diff(list[i], list[j]))
				}
			}
		}
	}

	return len(antinodes)
}

func diff(a, b space.Cell) space.Direction {
	return space.Direction{a[0] - b[0], a[1] - b[1]}
}

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 34)
	h.Expect(2, 9)
	h.Run()
}

//go:embed *.txt
var inputs embed.FS
