package main

import (
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
)

func solve(input [2][]string) int {
	var ranges [][]int
	for _, line := range input[0] {
		ranges = append(ranges, sti.Ints(strings.Split(line, "-")))
	}

	count := 0
	for _, index := range sti.Ints(input[1]) {
		for _, r := range ranges {
			if index >= r[0] && index <= r[1] {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 3)
	h.Run()
}
