package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	. "github.com/dbut2/advent-of-code/pkg/std"
)

func solve(input [][]string) int {
	count := 0
	for _, line := range input[len(input)-1] {
		ints := Ints(line)
		if ints[0]*ints[1] > 8*Sum(ints[2:]...) {
			count++
		}
	}
	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 2)
	h.Run()
}
