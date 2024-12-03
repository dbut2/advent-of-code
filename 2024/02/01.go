package main

import (
	"embed"

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
	h := harness.New(solve, input, tests)
	h.Tester.Expect(1, 2)
	h.Run()
}

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS
