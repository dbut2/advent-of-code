package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input []string) int {
	total := 0

	for _, line := range input {
		m := 0
		for i := 1; i < len(line); i++ {
			for j := 0; j < i; j++ {
				v := int(line[j]-'0')*10 + int(line[i]-'0')
				m = max(m, v)
			}
		}
		total += m
	}

	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 357)
	h.Run()
}
