package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
)

func solve(input []int) int {
	// vals holds the count of unique stones we have
	vals := map[int]int{}
	for _, n := range input {
		vals[n]++
	}

	// cache stones to their next stones
	cache := map[int][]int{0: {1}}

	for range 75 {
		next := map[int]int{}

		for n, amt := range vals {
			if _, ok := cache[n]; !ok {
				var entry []int
				if log10(n)%2 == 0 {
					ten := math.Pow(10, log10(n)/2)
					entry = []int{n / ten, n % ten}
				} else {
					entry = []int{n * 2024}
				}
				cache[n] = entry
			}

			for _, m := range cache[n] {
				next[m] += amt
			}
			continue
		}

		vals = next
	}

	c := 0
	for _, count := range vals {
		c += count
	}
	return c
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
	h.Run()
}
