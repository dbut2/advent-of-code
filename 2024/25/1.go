package main

import (
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input string) int {
	var keys [][5]int
	var locks [][5]int

	for _, in := range strings.Split(strings.TrimSpace(input), "\n\n") {
		heights := [5]int{}
		for _, line := range strings.Split(in, "\n") {
			for i, c := range line {
				if c == '#' {
					heights[i]++
				}
			}
		}
		if in[0] == '#' {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}

	count := 0
	for _, key := range keys {
		for _, lock := range locks {
			fits := true
			for i := range len(key) {
				if key[i]+lock[i] > 7 {
					fits = false
					break
				}
			}
			if fits {
				count++
			}
		}
	}
	return count
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 3)
	h.Run()
}
