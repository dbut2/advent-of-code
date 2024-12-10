package main

import (
	"slices"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
)

func solve(input [][]int) int {
	safeCount := 0
	for _, ints := range input {
		// brute force removing a single element for each line
		for i := range ints {
			if isSafe(slices.Concat(ints[:i], ints[i+1:])) {
				safeCount++
				break
			}
		}
	}

	return safeCount
}

func isSafe(ints []int) bool {
	shouldIncrease := ints[1] > ints[0]
	for i := 1; i < len(ints); i++ {
		distance := math.Abs(ints[i] - ints[i-1])
		if distance < 1 || distance > 3 {
			return false
		}
		increases := ints[i] > ints[i-1]
		if shouldIncrease != increases {
			return false
		}
	}
	return true
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 4)
	h.Run()
}
