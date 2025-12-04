package main

import (
	"time"

	"github.com/dbut2/advent-of-code/pkg/benchmark"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input space.Grid[byte]) int {
	count := 0

	q := lists.Queue[space.Cell]{}
	for cell := range input.Cells() {
		// start by checking all cells
		q.Push(cell)
	}

	for cell := range q.Seq {
		// check every cell in the queue
		v := input.Get(cell)
		if v == nil || *v != '@' {
			continue
		}
		c := 0
		for _, sur := range input.Surrounding(cell) {
			// if the surrounding cell is in the grid and a roll
			if sur != nil && *sur == '@' {
				c++
			}
		}
		if c < 4 {
			// remove by replacing roll char in grid and enqueue all neighbours
			*v = '.'
			count++
			for sur, _ := range input.Surrounding(cell) {
				q.Push(sur)
			}
		}
	}

	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 43)
	h.Benchmark(benchmark.Time(time.Second * 5))
	h.Run()
}
