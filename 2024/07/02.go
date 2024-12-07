package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
)

func solve(input [][]int) int {
	total := 0
	for _, line := range input {
		if calc(line[0], line[1], line[2:]) {
			total += line[0]
		}
	}
	return total
}

func calc(goal int, running int, nums []int) bool {
	if len(nums) == 0 {
		return goal == running
	}
	num, nums := nums[0], nums[1:]
	return calc(goal, running+num, nums) || calc(goal, running*num, nums) || calc(goal, concat(running, num), nums)
}

func concat(a, b int) int {
	return a*math.Pow(10, log10(b)) + b
}

func log10(a int) int {
	b := 0
	for a > 1 {
		b++
		a /= 10
	}
	return b
}

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 11387)
	h.Run()
}

//go:embed *.txt
var inputs embed.FS
