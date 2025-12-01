package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	. "github.com/dbut2/advent-of-code/pkg/std"
)

func solve(input []string) int {
	dial := 50
	count := 0

	for _, line := range input {
		for range Int(line[1:]) {
			switch line[0] {
			case 'L':
				dial--
			case 'R':
				dial++
			}
			if dial%100 == 0 {
				count++
			}
		}
	}

	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 6)
	h.Run()
}
