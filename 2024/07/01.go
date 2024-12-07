package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
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
	return calc(goal, running+num, nums) || calc(goal, running*num, nums)
}

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 3749)
	h.Run()
}

//go:embed *.txt
var inputs embed.FS
