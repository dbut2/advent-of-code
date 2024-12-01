package main

import (
	"embed"
	"slices"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/strings"
)

func solve(input []string) int {
	var left, right []int
	for _, line := range input {
		ints := strings.Ints(line)
		left = append(left, ints[0])
		right = append(right, ints[1])
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
	h := harness.New(solve, input, tests, harness.SplitNewlines())
	h.Run()
}

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS
