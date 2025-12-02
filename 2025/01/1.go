package main

import (
	"strconv"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input []string) int {
	dial, count := 50, 0
	for _, line := range input {
		amt, _ := strconv.Atoi(line[1:])
		dial += amt * (int(line[0]) - 79) / 3 // magic L=-1 R=1
		if dial%100 == 0 {
			count++
		}
	}
	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 3)
	h.Run()
}
