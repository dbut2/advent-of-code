package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input [][]int) int {
	var left, right []int
	for _, line := range input {
		left = append(left, line[0])
		right = append(right, line[1])
	}

	rCount := make(map[int]int)
	for _, r := range right {
		rCount[r]++
	}

	c := 0
	for _, l := range left {
		c += l * rCount[l]
	}
	return c
}

func main() {
	h := harness.New(solve, input, tests)
	h.Tester.Expect(1, 31)
	h.Run()
}

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS
