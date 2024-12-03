package main

import (
	"embed"
	"slices"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
)

func solve(input [][]int) int {
	var left, right []int
	for _, line := range input {
		left = append(left, line[0])
		right = append(right, line[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	c := 0
	for i := range left {
		c += math.Abs(left[i] - right[i])
	}
	return c
}

func main() {
	h := harness.New(solve, inputs)
	h.Run()
}

//go:embed *.txt
var inputs embed.FS
