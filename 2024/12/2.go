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
				upLeft := input.Get(cell.Move(dir).Move(dir.Rotate()))

				// Count corners

				// Case where cell is at a convex corner, both up and left are not equal
				if (up == nil || *up != *val) && (left == nil || *left != *val) {
					perimeter++
					continue
				}
				// Concave corner where left and up are equal but upLeft is not
				if up != nil && left != nil && *val == *up && *val == *left && *val != *upLeft {
					perimeter++
					continue
				}

				// All other cases of sides or internal cells
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
