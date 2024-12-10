package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input space.Grid[byte]) int {
	trailheads := input.FindAll(func(_ space.Cell, b byte) bool { return b == '0' })

	count := 0
	for _, trailhead := range trailheads {
		seen := sets.Set[space.Cell]{}

		queue := lists.Queue[space.Cell]{}
		queue.Push(trailhead)

		for cell := range queue.Seq {
			if seen.Contains(cell) {
				continue
			}
			seen.Add(cell)

			v := input.Get(cell)
			if *v == '9' {
				count++
				continue
			}

			for _, dir := range space.Directions {
				w := input.Get(cell.Move(dir))
				if w == nil {
					continue
				}

				if *w == *v+1 {
					queue.Push(cell.Move(dir))
				}
			}
		}
	}
	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 36)
	h.Run()
}
