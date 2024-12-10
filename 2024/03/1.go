package main

import (
	"regexp"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
)

func solve(input string) int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)

	total := 0
	for _, match := range matches {
		total += sti.Int(match[1]) * sti.Int(match[2])
	}

	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 161)
	h.Run()
}
