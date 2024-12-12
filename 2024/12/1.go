package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
)

type grid = space.Grid[byte]

func solve(input grid) int {
	all := sets.SetFrom(input.FindAll(func(_ space.Cell, _ byte) bool { return true }))

	areas := [][]space.Cell{}

	for len(all) > 0 {
		cell := all.Slice()[0]
		area := sets.Set[space.Cell]{}
		queue := lists.Queue[space.Cell]{cell}
		area.Add(cell)

		for cell := range queue.Seq {
			all.Remove(cell)
			val := input.Get(cell)
			for next, nextVal := range input.Adjacent(cell) {
				if area.Contains(next) {
					continue
				}
				if *val == *nextVal {
					area.Add(next)
					queue.Push(next)
				}
			}
		}

		areas = append(areas, area.Slice())
	}

	total := 0

	for _, area := range areas {
		perimeter := 0

		for _, cell := range area {
			for _, dir := range space.Directions {
				next := input.Get(cell.Move(dir))
				if next == nil {
					perimeter++
					continue
				}
				if *next != *input.Get(cell) {
					perimeter++
				}
			}
		}

		total += perimeter * len(area)
	}

	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 772)
	h.Expect(2, 1930)
	h.Run()
}
