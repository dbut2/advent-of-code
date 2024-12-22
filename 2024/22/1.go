package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input []int) int {
	total := 0
	for _, i := range input {
		for range 2000 {
			i ^= i * 64
			i &= 16777216 - 1
			i ^= i / 32
			i ^= i * 2048
			i &= 16777216 - 1
		}
		total += i
	}

	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 37327623)
	h.Run()
}
