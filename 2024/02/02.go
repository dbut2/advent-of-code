package main

import (
	"embed"
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
	increasing := ints[1] > ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] == ints[i-1] {
			return false
		}
		if increasing != (ints[i] > ints[i-1]) {
			return false
		}
		if math.Abs(ints[i]-ints[i-1]) > 3 {
			return false
		}
	}
	return true
}

func main() {
	h := harness.New(solve, input, tests, harness.SplitNewlinesWithInts())
	h.Tester.Expect(1, 4)
	h.Run()
}

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS
