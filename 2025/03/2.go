package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input []string) int {
	total := 0

	for _, line := range input {
		v := 0
		// greedy search for largest digit while keeping enough remaining for
		// rest of number
		for i := range 12 {
			m := rune(0)
			n := 0
			for j, c := range line[:len(line)-(11-i)] {
				if c > m {
					m = c
					n = j
				}
			}
			v *= 10
			v += int(m) - '0'
			line = line[n+1:]
		}
		total += v
	}

	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 3121910778619)
	h.Run()
}
