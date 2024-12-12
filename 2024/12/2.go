package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input space.Grid[byte]) int {
	floods := space.FloodAll(input)

	total := 0
	for _, flood := range floods {
		perimeter := 0
		for _, cell := range flood {
			for _, dir := range space.Directions {
				val := input.Get(cell)
				left := input.Get(cell.Move(dir))
				up := input.Get(cell.Move(dir.Rotate()))
				leftUp := input.Get(cell.Move(dir).Move(dir.Rotate()))

				// These combinations of states are what I drew out in pen and
				// paper. I cannot explain further why.
				if (up == nil || *up != *val) && (left == nil || *left != *val) {
					perimeter++
					continue
				}
				if up == nil || left == nil {
					continue
				}
				if *val == *up && *val != *left && *val == *leftUp {
					perimeter++
					continue
				}
			}
		}
		total += perimeter * len(flood)
	}
	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(3, 80)
	h.Expect(1, 436)
	h.Expect(2, 1206)
	h.Run()
}
