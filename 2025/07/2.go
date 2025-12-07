package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input []string) int {
	prevLineCount := make([]int, len(input[0]))
	for i, char := range input[0] {
		if char == 'S' {
			prevLineCount[i] = 1
		}
	}

	for _, line := range input {
		lineCount := make([]int, len(line))
		for i, char := range line {
			if char == '^' {
				lineCount[i-1] += prevLineCount[i]
				lineCount[i+1] += prevLineCount[i]
			} else {
				lineCount[i] += prevLineCount[i]
			}
		}
		prevLineCount = lineCount
	}

	total := 0
	for _, count := range prevLineCount {
		total += count
	}
	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 40)
	h.Run()
}
