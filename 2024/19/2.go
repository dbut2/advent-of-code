package main

import (
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input []string) int {
	supply := strings.Split(input[0], ", ")
	cache := map[string]int{"": 1}
	var count func(string) int
	count = func(towel string) int {
		if v, ok := cache[towel]; ok {
			return v
		}
		total := 0
		for _, s := range supply {
			if strings.HasPrefix(towel, s) {
				total += count(strings.TrimPrefix(towel, s))
			}
		}
		cache[towel] = total
		return total
	}
	total := 0
	for _, towel := range input[2:] {
		total += count(towel)
	}
	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 16)
	h.Run()
}
