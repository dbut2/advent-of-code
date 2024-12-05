package main

import (
	"embed"
	"slices"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/sti"
)

func solve(input string) int {
	sections := strings.Split(input, "\n\n")
	orderingRules, updates := sections[0], sections[1]

	// Create a dependancy map of Key depends on []Values
	dep := make(map[string][]string)

	for _, line := range strings.Split(orderingRules, "\n") {
		parts := strings.Split(line, "|")
		depender, dependant := parts[0], parts[1]
		dep[dependant] = append(dep[dependant], depender)
	}

	total := 0

	for _, line := range strings.Split(updates, "\n") {
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
			// I am ashamed of my solution here, but I will post anyway as
			// documentation of my solution.

			// create a set of all elements in the
			lineElements := sets.SetFrom(parts)

			// m is a secondary dependancy map, containing only keys and values
			// that exist in the current update.
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
			//the element that has no dependencies remaining, meaning this must
			//be the next value in our sorted string.
			var s []string
			for len(m) > 0 {
				for _, p := range lineElements.Slice() {
					if len(m[p]) == 0 {
						// Add current element to end of slice.
						s = append(s, p)

						// Remove current element from dependency map.
						for i := range m {
							m[i] = slices.DeleteFunc(m[i], func(s string) bool {
								return s == p
							})
						}
						delete(m, p)
						lineElements.Remove(p)

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
	h := harness.New(solve, inputs)
	h.Expect(1, 123)
	h.Run()
}

//go:embed *.txt
var inputs embed.FS
