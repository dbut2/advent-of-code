package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input [][]int) int {
	tokens := 0

	for i := range len(input) / 4 {
		a0, a1 := input[4*i][0], input[4*i][1]
		b0, b1 := input[4*i+1][0], input[4*i+1][1]
		p0, p1 := input[4*i+2][0], input[4*i+2][1]

		aNum := b0*p1 - p0*b1
		aDen := b0*a1 - a0*b1
		bNum := a0*p1 - p0*a1
		bDem := a0*b1 - b0*a1

		if aNum%aDen == 0 && bNum%bDem == 0 {
			tokens += 3*aNum/aDen + bNum/bDem
		}
	}

	return tokens
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 480)
	h.Run()
}
