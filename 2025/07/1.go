package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
	. "github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input Grid[byte]) int {
	start := input.Find(func(b byte) bool {
		return b == 'S'
	})

	count := 0

	q := lists.NewQueue[Cell]()
	q.Push(start)

	seen := sets.Set[Cell]{}
	for cell := range q.Seq {
		if seen.Contains(cell) {
			continue
		}
		seen.Add(cell)

		next := cell.Move(Down)
		if !input.Inside(next) {
			continue
		}

		if *input.Get(next) == '^' {
			count++
			q.Push(next.Move(Left))
			q.Push(next.Move(Right))
		} else {
			q.Push(next)
		}
	}

	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 21)
	h.Run()
}
