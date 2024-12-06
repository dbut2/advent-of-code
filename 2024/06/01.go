package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input space.Grid[byte]) int {
	cell, _ := input.Find(func(cell space.Cell, b byte) bool { return b == '^' })
	dir := space.Up

	seen := sets.Set[space.Cell]{}
	for {
		seen.Add(cell)

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

	return len(seen)
}

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 41)
	h.Run()
}

//go:embed *.txt
var inputs embed.FS
