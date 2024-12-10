package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
)

func solve(input [][]int) int {
	safeCount := 0
	for _, ints := range input {
		increasing := ints[1] > ints[0]
		isSafe := true
		for i := 1; i < len(ints); i++ {
			if ints[i] == ints[i-1] {
				isSafe = false
				break
			}
			if increasing != (ints[i] > ints[i-1]) {
				isSafe = false
				break
			}
			if math.Abs(ints[i]-ints[i-1]) > 3 {
				isSafe = false
				break
			}
		}
		if isSafe {
			safeCount++
		}
	}

	return safeCount
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 2)
	h.Run()
}
