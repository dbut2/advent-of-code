package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
)

func solve(input []int) int {
	for range 25 {
		next := []int{}
		for i := range input {
			c := input[i]
			if c == 0 {
				next = append(next, 1)
			} else if log10(c)%2 == 0 {
				ten := math.Pow(10, log10(c)/2)
				next = append(next, c/ten, c%ten)
			} else {
				next = append(next, c*2024)
			}

		}
		input = next
	}
	return len(input)
}

func log10(n int) int {
	i := 1
	for n >= 10 {
		i++
		n /= 10
	}
	return i
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 55312)
	h.Run()
}
