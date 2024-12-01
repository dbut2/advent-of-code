package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/strings"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 31)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var left []int
	var right []int

	for j, line := range s {
		_, _ = j, line

		pair := strings.Ints(line)
		left = append(left, pair[0])
		right = append(right, pair[1])
	}

	rCount := make(map[int]int)
	for _, r := range right {
		rCount[r] += 1
	}

	c := 0
	for _, l := range left {
		c += l * rCount[l]
	}
	return c
}
