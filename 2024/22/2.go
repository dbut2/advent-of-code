package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input []int) int {
	key := func(a, b, c, d int) int {
		k := (a + 9) * 19
		k += b + 9
		k *= 19
		k += c + 9
		k *= 19
		k += d + 9
		return k
	}

	counts := [19 * 19 * 19 * 19]int{}

	for _, i := range input {
		var a, b, c, d int
		prev := i
		seen := [19 * 19 * 19 * 19]bool{}
		for k := range 2000 {
			i ^= i * 64
			i &= 16777216 - 1
			i ^= i / 32
			i ^= i * 2048
			i &= 16777216 - 1

			a, b, c, d = b, c, d, i%10-prev%10
			prev = i

			if k >= 4 {
				if seen[key(a, b, c, d)] {
					continue
				}
				seen[key(a, b, c, d)] = true
				counts[key(a, b, c, d)] += i % 10
			}
		}
	}

	m := 0
	for _, v := range counts {
		m = max(m, v)
	}
	return m
}

func main() {
	h := harness.New(solve)
	h.Expect(2, 23)
	h.Run()
}
