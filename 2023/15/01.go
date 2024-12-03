package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Tester.Expect(1, 1320)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input, ",")

	total := 0
	for _, line := range s {
		hash := 0
		for _, char := range line {
			hash += int(char)
			hash *= 17
			hash %= 256
		}
		total += hash
	}
	return total
}
