package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	. "github.com/dbut2/advent-of-code/pkg/std"
)

func solve(input [][]int) int {
	maxArea := 0
	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			a, b := input[i], input[j]
			area := (Abs(a[0]-b[0]) + 1) * (Abs(a[1]-b[1]) + 1)
			maxArea = max(maxArea, area)
		}
	}
	return maxArea
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 50)
	h.Run()
}
