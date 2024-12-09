package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input space.Grid[byte]) int {
	count := 0

	for c, val := range input.Cells() {
		pre := *val
		if pre == '^' {
			continue
		}
		input.Set(c, '#')

		cell, _ := input.Find(func(_ space.Cell, b byte) bool { return b == '^' })
		dir := space.Up

		type pos struct {
			cell space.Cell
			dir  space.Direction
		}
		seen := make(sets.Set[pos], len(input)*len(input[0]))
		for {
			p := pos{cell: cell, dir: dir}
			if seen.Contains(p) {
				count++
				break
			}
			seen.Add(p)

			next := input.Get(cell.Move(dir))
			if next == nil {
				break
			}

			if *next == '#' {
				dir = dir.Rotate()
			} else {
				cell = cell.Move(dir)
			}
		}

		input.Set(c, pre)
	}

	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 6)
	h.Run()
}
