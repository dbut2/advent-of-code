package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 21)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0
	for _, line := range s {
		springs := strings.Split(line, " ")[0]
		goals := strings.Split(line, " ")[1]

		goalNumbers := sti.Stis(strings.Split(goals, ","))

		total += validSubsets(springs, goalNumbers)
	}
	return total
}

func validSubsets(springs string, goalNumbers []int) int {
	count := 0

	if !strings.Contains(springs, "?") {
		if valid(springs, goalNumbers) {
			return 1
		} else {
			return 0
		}
	}

	a := strings.Replace(springs, "?", "#", 1)
	b := strings.Replace(springs, "?", ".", 1)

	count += validSubsets(a, goalNumbers)
	count += validSubsets(b, goalNumbers)

	return count
}

func valid(springs string, goalNumbers []int) bool {
	realNumbers := []int{}
	count := 0
	for _, char := range springs {
		switch char {
		case '#':
			count++
		case '.':
			if count != 0 {
				realNumbers = append(realNumbers, count)
			}
			count = 0
		}
	}
	if count != 0 {
		realNumbers = append(realNumbers, count)
	}

	if len(realNumbers) != len(goalNumbers) {
		return false
	}

	for i := range realNumbers {
		if realNumbers[i] != goalNumbers[i] {
			return false
		}
	}

	return true
}
