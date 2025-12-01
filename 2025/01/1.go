package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	. "github.com/dbut2/advent-of-code/pkg/std"
)

func solve(input []string) int {
	dial := 50
	count := 0

	for _, line := range input {
		switch line[0] {
		case 'L':
			dial -= Int(line[1:])
		case 'R':
			dial += Int(line[1:])
		}
		dial %= 100
		if dial == 0 {
			count++
		}
	}

	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 3)
	h.Run()
}
