package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input []string) int {
	prevLine := make([]int, len(input[0]))
	for x, v := range input[0] {
		if v == 'S' {
			prevLine[x] = 1
		}
	}

	for _, line := range input {
		lineCount := make([]int, len(line))
		for j, char := range line {
			if char == '^' {
				lineCount[j-1] += prevLine[j]
				lineCount[j+1] += prevLine[j]
			} else {
				lineCount[j] += prevLine[j]
			}
		}
		prevLine = lineCount
	}

	count := 0
	for _, v := range prevLine {
		count += v
	}
	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 40)
	h.Run()
}
