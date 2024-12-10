package main

import (
	"slices"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/sti"
)

func solve(input [2][]string) int {
	orderingRules, updates := input[0], input[1]

	// Create a dependancy map of Key depends on []Values
	dep := make(map[string][]string)

	for _, line := range orderingRules {
		parts := strings.Split(line, "|")
		depender, dependant := parts[0], parts[1]
		dep[dependant] = append(dep[dependant], depender)
	}

	total := 0

	for _, line := range updates {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")

		valid := true
		for i, part := range parts {
			for j := i + 1; j < len(parts); j++ {
				// If there exists a later value that the current value depends
				// on, this is invalid as it needs to be to the left.
				if slices.Contains(dep[part], parts[j]) {
					valid = false
				}
			}
		}

		if !valid {
			// I am ashamed of this solution.

			// Create a set of all elements in the update.
			lineElements := sets.SetFrom(parts)

			// Create secondary dependency map. A copy of the original
			// dependency map, only including keys and values that exist in the
			// current update.
			m := make(map[string][]string)
			for x, y := range dep {
				if !lineElements.Contains(x) {
					continue
				}
				for _, p := range y {
					if lineElements.Contains(p) {
						m[x] = append(m[x], p)
					}
				}
			}

			// While there exists some elements in the dependency map, pull out
			// the element that has no dependencies remaining, this must be the
			// next value in the sequence.
			var s []string
			for len(m) > 0 {
				for _, p := range lineElements.Slice() {
					if len(m[p]) == 0 {
						// Add element to end of slice.
						s = append(s, p)

						// Remove references to element.
						lineElements.Remove(p)
						delete(m, p)
						for i := range m {
							m[i] = slices.DeleteFunc(m[i], func(s string) bool {
								return s == p
							})
						}

						break
					}
				}
			}

			total += sti.Int(s[len(s)/2])
		}
	}

	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 123)
	h.Run()
}
