package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/std"
)

func solve(input space.Grid[byte]) int {
	start := input.Find(std.Equal[byte]('S'))
	end := input.Find(std.Equal[byte]('E'))

	// Create a map of all cells in the path from start to end with the values
	// of their distance from start
	times := make(map[space.Cell]int)
	curr := start
	for {
		times[curr] = len(times)
		if curr == end {
			break
		}
		for next, v := range input.Adjacent(curr) {
			if *v == '#' {
				continue
			}
			if _, ok := times[next]; ok {
				continue
			}
			curr = next
		}
	}

	count := 0
	for cell, v := range input.Cells() {
		if *v == '#' {
			continue
		}
		const k = 20
		for dx := -k; dx <= k; dx++ {
			for dy := -k; dy <= k; dy++ {
				// For all cells in the surrounding 2k+1 * 2k+1 area, find all
				// cells with a manhattan distance of up to k.
				// For those, cells if the benefit of cheating would result in
				// skipping 100 or greater cells minus the manhattan distance
				// cost, increase count.

				next := cell.Move(space.Direction{dx, dy})
				if space.Manhattan(cell, next) > k {
					continue
				}
				w := input.Get(next)
				if w == nil || *w == '#' {
					continue
				}
				diff := times[next] - times[cell] - space.Manhattan(cell, next)
				if diff >= 100 {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 0)
	h.Run()
}
